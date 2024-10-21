package interceptor

import (
	"connectrpc.com/validate"
)

type ValidationInterceptorAdapter interface {
	ValidationInterceptor() *validate.Interceptor
}

type ValidationInterceptorImpl struct {
}

func NewValidationInterceptorAdapter() ValidationInterceptorAdapter {
	return &ValidationInterceptorImpl{}
}

func (l *ValidationInterceptorImpl) ValidationInterceptor() *validate.Interceptor {
	interceptor, err := validate.NewInterceptor()
	if err != nil {
		panic(err)
	}
	return interceptor
}
