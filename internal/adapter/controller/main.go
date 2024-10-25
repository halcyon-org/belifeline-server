package controller

import (
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
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
	extinfo  mainv1connect.ExternalInformationServiceHandler
	koyo     mainv1connect.KoyoServiceHandler
	health   mainv1connect.HealthServiceHandler

	auth    interceptor.AuthInterceptorAdapter
	logging interceptor.LoggingInterceptorAdapter

	cfg config.Repository
}

func NewBeLifelineController(
	admin mainv1connect.AdminServiceHandler,
	provider mainv1connect.ProviderServiceHandler,
	extinfo mainv1connect.ExternalInformationServiceHandler,
	koyo mainv1connect.KoyoServiceHandler,
	health mainv1connect.HealthServiceHandler,
	auth interceptor.AuthInterceptorAdapter,
	logging interceptor.LoggingInterceptorAdapter,
	config config.Repository,
) BeLifelineController {
	return &BeLifelineControllerImpl{
		admin:    admin,
		provider: provider,
		extinfo:  extinfo,
		koyo:     koyo,
		health:   health,
		auth:     auth,
		logging:  logging,
		cfg:      config,
	}
}

func (c *BeLifelineControllerImpl) Serve() error {
	err := c.cfg.LoadConfig()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector("belifeline.v1")
	mux.Handle(grpcreflect.NewHandlerV1(reflector, connect.WithInterceptors(c.logging.LoggingInterceptor())))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector, connect.WithInterceptors(c.logging.LoggingInterceptor())))
	mux.Handle(mainv1connect.NewAdminServiceHandler(
		c.admin,
		connect.WithInterceptors(c.logging.LoggingInterceptor(), c.auth.AuthAdminServiceInterceptor()),
	))
	mux.Handle(mainv1connect.NewProviderServiceHandler(
		c.provider,
		connect.WithInterceptors(c.logging.LoggingInterceptor(), c.auth.AuthProviderServiceInterceptor()),
	))
	mux.Handle(mainv1connect.NewExternalInformationServiceHandler(
		c.extinfo,
		connect.WithInterceptors(c.logging.LoggingInterceptor(), c.auth.AuthExternalInformationServiceInterceptor()),
	))
	mux.Handle(mainv1connect.NewKoyoServiceHandler(
		c.koyo,
		connect.WithInterceptors(c.logging.LoggingInterceptor(), c.auth.AuthKoyoServiceInterceptor()),
	))
	mux.Handle(mainv1connect.NewHealthServiceHandler(
		c.health,
		connect.WithInterceptors(c.logging.LoggingInterceptor()),
	))

	servAddr := fmt.Sprintf("%s:%d", c.cfg.GetConfig().ListenAddr, c.cfg.GetConfig().Port)
	log.Printf("Listening on %s\n", servAddr)
	err = http.ListenAndServe(
		servAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

	return err
}
