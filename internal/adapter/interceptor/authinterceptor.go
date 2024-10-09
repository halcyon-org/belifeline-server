package interceptor

import (
	"context"
	"fmt"

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

const AuthAPIKeyHeader = "X-API-Key"

type AuthHeaderType string

const (
	AdminUserKey  AuthHeaderType = "admin_user"
	ClientDataKey AuthHeaderType = "client_data"
	KoyoInfoKey   AuthHeaderType = "koyo_info"
	ExtInfoKey    AuthHeaderType = "ext_info"
)

type AuthInterceptorAdapter interface {
	AuthAdminServiceInterceptor() connect.UnaryInterceptorFunc
	AuthProviderServiceInterceptor() connect.UnaryInterceptorFunc
}

type AuthInterceptorImpl struct {
	authUsecase usecase.AuthUsecase
}

var ErrMissingAPIKey = fmt.Errorf("missing %s header", AuthAPIKeyHeader)

func NewAuthInterceptorAdapter(authUsecase usecase.AuthUsecase) AuthInterceptorAdapter {
	return &AuthInterceptorImpl{
		authUsecase: authUsecase,
	}
}

func (a *AuthInterceptorImpl) AuthAdminServiceInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				return next(ctx, req)
			}

			apiKey := req.Header().Get(AuthAPIKeyHeader)
			if apiKey == "" {
				return nil, connect.NewError(connect.CodePermissionDenied, ErrMissingAPIKey)
			}

			user, err := a.authUsecase.AuthAdminUser(ctx, apiKey)
			if err != nil {
				return nil, connect.NewError(connect.CodePermissionDenied, err)
			}
			ctx = context.WithValue(ctx, AdminUserKey, user)
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

func (a *AuthInterceptorImpl) AuthProviderServiceInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				return next(ctx, req)
			}

			apiKey := req.Header().Get(AuthAPIKeyHeader)
			if apiKey == "" {
				return nil, connect.NewError(connect.CodePermissionDenied, ErrMissingAPIKey)
			}

			admin, err := a.authUsecase.AuthAdminUser(ctx, apiKey)
			if err == nil && admin != nil {
				ctx = context.WithValue(ctx, AdminUserKey, admin)
				return next(ctx, req)
			}

			data, err := a.authUsecase.AuthClientData(ctx, apiKey)
			if err == nil && data != nil {
				ctx = context.WithValue(ctx, ClientDataKey, data)
				return next(ctx, req)
			}

			koyo, err := a.authUsecase.AuthKoyoInformation(ctx, apiKey)
			if err == nil && koyo != nil {
				ctx = context.WithValue(ctx, KoyoInfoKey, koyo)
				return next(ctx, req)
			}

			extinfo, err := a.authUsecase.AuthExternalInformation(ctx, apiKey)
			if err == nil && extinfo != nil {
				ctx = context.WithValue(ctx, ExtInfoKey, extinfo)
				return next(ctx, req)
			}

			return nil, connect.NewError(connect.CodePermissionDenied, err)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
