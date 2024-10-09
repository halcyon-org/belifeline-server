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
	admin    mainv1connect.AdminServiceHandler
	provider mainv1connect.ProviderServiceHandler
	extinfo  mainv1connect.ExtInfoServiceHandler
	koyo     mainv1connect.KoyoServiceHandler
	server   mainv1connect.ServerServiceHandler

	auth interceptor.AuthInterceptorAdapter

	cfg config.ConfigRepository
}

func NewBeLifelineController(admin mainv1connect.AdminServiceHandler, provider mainv1connect.ProviderServiceHandler,
	extinfo mainv1connect.ExtInfoServiceHandler,
	koyo mainv1connect.KoyoServiceHandler,
	server mainv1connect.ServerServiceHandler, auth interceptor.AuthInterceptorAdapter, config config.ConfigRepository) BeLifelineController {
	return &BeLifelineControllerImpl{
		admin:    admin,
		provider: provider,
		extinfo:  extinfo,
		koyo:     koyo,
		server:   server,
		auth:     auth,
		cfg:      config,
	}
}

func (c *BeLifelineControllerImpl) Serve() error {
	err := c.cfg.LoadConfig()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	interceptor := connect.WithInterceptors(c.auth.AuthInterceptor())
	mux.Handle(mainv1connect.NewAdminServiceHandler(c.admin, interceptor))
	mux.Handle(mainv1connect.NewProviderServiceHandler(c.provider, interceptor))
	mux.Handle(mainv1connect.NewExtInfoServiceHandler(c.extinfo, interceptor))
	mux.Handle(mainv1connect.NewKoyoServiceHandler(c.koyo, interceptor))
	mux.Handle(mainv1connect.NewServerServiceHandler(c.server, interceptor))

	err = http.ListenAndServe(
		fmt.Sprintf("%s:%d", c.cfg.GetConfig().ListenAddr, c.cfg.GetConfig().Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

	return err
}
