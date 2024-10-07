package controller

import (
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"github.com/halcyon-org/kizuna/internal/adapter/interceptor"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type BeLifelineController interface {
	Serve() error
}

type BeLifelineControllerImpl struct {
	server mainv1connect.BeLifelineServiceHandler
	auth   interceptor.AuthInterceptorAdapter

	cfg config.ConfigRepository
}

func NewBeLifelineController(server mainv1connect.BeLifelineServiceHandler, auth interceptor.AuthInterceptorAdapter, config config.ConfigRepository) BeLifelineController {
	return &BeLifelineControllerImpl{
		server: server,
		auth:   auth,
		cfg:    config,
	}
}

func (c *BeLifelineControllerImpl) Serve() error {
	err := c.cfg.LoadConfig()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	interceptor := connect.WithInterceptors(c.auth.AuthInterceptor())
	path, handler := mainv1connect.NewBeLifelineServiceHandler(c.server, interceptor)
	mux.Handle(path, handler)

	err = http.ListenAndServe(
		fmt.Sprintf("%s:%d", c.cfg.GetConfig().ListenAddr, c.cfg.GetConfig().Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

	return err
}
