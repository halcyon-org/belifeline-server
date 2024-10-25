package interceptor

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/halcyon-org/kizuna/internal/usecase"
)

const (
	authTypeNoAuth = iota
	authTypeAdmin
	authTypeClient
	authTypeKoyo
	authTypeExitInfo
)

const AuthAPIHeader = "X-API-Key"

type AuthHeaderType string

const (
	AdminUserKey           AuthHeaderType = "admin_user"
	ClientInformationKey   AuthHeaderType = "client_data"
	KoyoInfoKey            AuthHeaderType = "koyo_info"
	ExternalInformationKey AuthHeaderType = "ext_info"
)

type AuthInterceptorAdapter interface {
	AuthAdminServiceInterceptor() connect.UnaryInterceptorFunc
	AuthProviderServiceInterceptor() connect.UnaryInterceptorFunc
	AuthKoyoServiceInterceptor() connect.UnaryInterceptorFunc
	AuthExternalInformationServiceInterceptor() connect.UnaryInterceptorFunc
}

type AuthInterceptorImpl struct {
	authUsecase usecase.AuthUsecase
}

var ErrMissingAPIKey = fmt.Errorf("missing %s header", AuthAPIHeader)

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

			apiKey := req.Header().Get(AuthAPIHeader)
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

			apiKey := req.Header().Get(AuthAPIHeader)
			if apiKey == "" {
				return nil, connect.NewError(connect.CodePermissionDenied, ErrMissingAPIKey)
			}

			client, err := a.authUsecase.AuthClientInformation(ctx, apiKey)
			if err != nil {
				return nil, connect.NewError(connect.CodePermissionDenied, err)
			}

			ctx = context.WithValue(ctx, ClientInformationKey, client)

			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}

func (a *AuthInterceptorImpl) AuthKoyoServiceInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				return next(ctx, req)
			}

			apiKey := req.Header().Get(AuthAPIHeader)
			if apiKey == "" {
				return nil, connect.NewError(connect.CodePermissionDenied, ErrMissingAPIKey)
			}

			koyo, err := a.authUsecase.AuthKoyoInformation(ctx, apiKey)
			if err != nil {
				return nil, connect.NewError(connect.CodePermissionDenied, err)
			}

			ctx = context.WithValue(ctx, KoyoInfoKey, koyo)

			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}

func (a *AuthInterceptorImpl) AuthExternalInformationServiceInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				return next(ctx, req)
			}

			apiKey := req.Header().Get(AuthAPIHeader)
			if apiKey == "" {
				return nil, connect.NewError(connect.CodePermissionDenied, ErrMissingAPIKey)
			}

			ext, err := a.authUsecase.AuthExternalInformation(ctx, apiKey)
			if err != nil {
				return nil, connect.NewError(connect.CodePermissionDenied, err)
			}

			ctx = context.WithValue(ctx, ExternalInformationKey, ext)

			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}
