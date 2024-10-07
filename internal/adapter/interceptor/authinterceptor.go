package interceptor

import (
	"context"
	"fmt"
	"slices"

	"connectrpc.com/connect"
	"github.com/halcyon-org/kizuna/internal/usecase"
)

const (
	AUTH_TYPE_NO_AUTH = iota
	AUTH_TYPE_ADMIN
	AUTH_TYPE_CLIENT
	AUTH_TYPE_KOYO
	AUTH_TYPE_EXTINFO
)

var authTypeByMethod = map[string][]int{
	"/belifeline.v1.BeLifelineService/ClientCreate":  {AUTH_TYPE_ADMIN},
	"/belifeline.v1.BeLifelineService/ClientList":    {AUTH_TYPE_ADMIN},
	"/belifeline.v1.BeLifelineService/ClientDelete":  {AUTH_TYPE_ADMIN},
	"/belifeline.v1.BeLifelineService/ClientRevoke":  {AUTH_TYPE_ADMIN},
	"/belifeline.v1.BeLifelineService/ExtInfoCreate": {AUTH_TYPE_ADMIN},
	"/belifeline.v1.BeLifelineService/ExtInfoList":   {AUTH_TYPE_ADMIN, AUTH_TYPE_CLIENT, AUTH_TYPE_KOYO, AUTH_TYPE_KOYO},
	"/belifeline.v1.BeLifelineService/ExtInfoDelete": {AUTH_TYPE_ADMIN},
	"/belifeline.v1.BeLifelineService/KoyoCreate":    {AUTH_TYPE_ADMIN},
	"/belifeline.v1.BeLifelineService/KoyoList":      {AUTH_TYPE_ADMIN, AUTH_TYPE_CLIENT, AUTH_TYPE_KOYO},
	"/belifeline.v1.BeLifelineService/KoyoDelete":    {AUTH_TYPE_ADMIN},
	"/belifeline.v1.BeLifelineService/KoyoApiRevoke": {AUTH_TYPE_ADMIN},
	"/belifeline.v1.BeLifelineService/Health":        {AUTH_TYPE_NO_AUTH},
}

const AuthAPIKeyHeader = "X-API-Key"

type AuthHeaderType string

const (
	AdminUserKey AuthHeaderType = "admin_user"
)

type AuthInterceptorAdapter interface {
	AuthInterceptor() connect.UnaryInterceptorFunc
}

type AuthInterceptorImpl struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthInterceptorAdapter(authUsecase usecase.AuthUsecase) AuthInterceptorAdapter {
	return &AuthInterceptorImpl{
		authUsecase: authUsecase,
	}
}

func (a *AuthInterceptorImpl) AuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				return next(ctx, req)
			}

			val, ok := authTypeByMethod[req.Spec().Procedure]
			if !ok {
				return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("method %s not supported", req.Spec().Procedure))
			}

			if slices.Contains(val, AUTH_TYPE_NO_AUTH) {
				return next(ctx, req)
			}

			apiKey := req.Header().Get(AuthAPIKeyHeader)
			if apiKey == "" {
				return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("missing %s header", AuthAPIKeyHeader))
			}

			if slices.Contains(val, AUTH_TYPE_ADMIN) {
				user, err := a.authUsecase.AuthAdminUser(ctx, apiKey)
				if err != nil {
					return nil, connect.NewError(connect.CodePermissionDenied, err)
				}
				ctx = context.WithValue(ctx, AdminUserKey, user)
				return next(ctx, req)
			}

			return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("method %s not supported", req.Spec().Procedure))
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
