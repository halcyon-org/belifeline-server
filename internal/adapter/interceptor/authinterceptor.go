package interceptor

import (
	"context"

	"connectrpc.com/connect"
)

const (
	AUTH_TYPE_NO_AUTH = iota
	AUTH_TYPE_ADMIN
	AUTH_TYPE_CLIENT
	AUTH_TYPE_KOYO
	AUTH_TYPE_EXTINFO
)

var authTypeByMethod = map[string][]int{
	"/main.v1.BeLifelineService/ClientCreate":  {AUTH_TYPE_ADMIN},
	"/main.v1.BeLifelineService/ClientList":    {AUTH_TYPE_ADMIN},
	"/main.v1.BeLifelineService/ClientDelete":  {AUTH_TYPE_ADMIN},
	"/main.v1.BeLifelineService/ClientRevoke":  {AUTH_TYPE_ADMIN},
	"/main.v1.BeLifelineService/ExtInfoCreate": {AUTH_TYPE_ADMIN},
	"/main.v1.BeLifelineService/ExtInfoList":   {AUTH_TYPE_ADMIN, AUTH_TYPE_CLIENT, AUTH_TYPE_KOYO, AUTH_TYPE_KOYO},
	"/main.v1.BeLifelineService/ExtInfoDelete": {AUTH_TYPE_ADMIN},
	"/main.v1.BeLifelineService/KoyoCreate":    {AUTH_TYPE_ADMIN},
	"/main.v1.BeLifelineService/KoyoList":      {AUTH_TYPE_ADMIN, AUTH_TYPE_CLIENT, AUTH_TYPE_KOYO},
	"/main.v1.BeLifelineService/KoyoDelete":    {AUTH_TYPE_ADMIN},
	"/main.v1.BeLifelineService/KoyoApiRevoke": {AUTH_TYPE_ADMIN},
	"/main.v1.BeLifelineService/Health":        {AUTH_TYPE_NO_AUTH},
}

type AuthInterceptorAdapter interface {
	AuthInterceptor() connect.UnaryInterceptorFunc
}

type AuthInterceptorImpl struct {
}

func NewAuthInterceptorAdapter() AuthInterceptorAdapter {
	return &AuthInterceptorImpl{}
}

func (a *AuthInterceptorImpl) AuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
