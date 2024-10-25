package interceptor

import (
	"context"
	"log/slog"
	"os"

	"connectrpc.com/connect"
	"github.com/google/uuid"
)

type ContextKey string

const (
	LoggerKey        ContextKey = "Logger"
	HeaderResponseID string     = "ResponseID"
)

type LoggingInterceptorAdapter interface {
	LoggingInterceptor() connect.UnaryInterceptorFunc
}

type LoggingInterceptorImpl struct{}

func NewLoggingInterceptorAdapter() LoggingInterceptorAdapter {
	return &LoggingInterceptorImpl{}
}

func (l *LoggingInterceptorImpl) LoggingInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			resID := uuid.New().String()

			logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With(HeaderResponseID, resID)

			nctx := context.WithValue(ctx, LoggerKey, logger)

			logger.Info(
				"Request",
				"Procedure", req.Spec().Procedure,
				"Protocol", req.Peer().Protocol,
				"Addr", req.Peer().Addr,
			)

			res, err := next(nctx, req)
			if res == nil {
				return res, err
			}

			res.Header().Set(HeaderResponseID, resID)

			return res, err
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}
