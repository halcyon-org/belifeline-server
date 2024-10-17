package interceptor

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
)

type LoggingInterceptorAdapter interface {
	LoggingInterceptor() connect.UnaryInterceptorFunc
}

type LoggingInterceptorImpl struct {
}

func NewLoggingInterceptorAdapter() LoggingInterceptorAdapter {
	return &LoggingInterceptorImpl{}
}

func (l *LoggingInterceptorImpl) LoggingInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			slog.Info(
				"Request",
				"Procedure", req.Spec().Procedure,
				"Protocol", req.Peer().Protocol,
				"Addr", req.Peer().Addr,
			)
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
