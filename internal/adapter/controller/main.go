package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"github.com/halcyon-org/kizuna/internal/adapter/interceptor"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type BeLifelineController interface {
	makeMux() *http.ServeMux

	Serve() error
	TestServe() (*httptest.Server, error)
}

type BeLifelineControllerImpl struct {
	admin    mainv1connect.AdminServiceHandler
	provider mainv1connect.ProviderServiceHandler
	extinfo  mainv1connect.ExternalInformationServiceHandler
	koyo     mainv1connect.KoyoServiceHandler
	health   mainv1connect.HealthServiceHandler

	auth       interceptor.AuthInterceptorAdapter
	logging    interceptor.LoggingInterceptorAdapter
	validation interceptor.ValidationInterceptorAdapter

	cfg config.ConfigRepository
}

func NewBeLifelineController(admin mainv1connect.AdminServiceHandler, provider mainv1connect.ProviderServiceHandler,
	extinfo mainv1connect.ExternalInformationServiceHandler,
	koyo mainv1connect.KoyoServiceHandler,
	health mainv1connect.HealthServiceHandler, auth interceptor.AuthInterceptorAdapter, logging interceptor.LoggingInterceptorAdapter, validation interceptor.ValidationInterceptorAdapter, config config.ConfigRepository) BeLifelineController {
	return &BeLifelineControllerImpl{
		admin:      admin,
		provider:   provider,
		extinfo:    extinfo,
		koyo:       koyo,
		health:     health,
		auth:       auth,
		logging:    logging,
		validation: validation,
		cfg:        config,
	}
}

func (c *BeLifelineControllerImpl) makeMux() *http.ServeMux {
	mux := http.NewServeMux()

	reflector := grpcreflect.NewStaticReflector("belifeline.v1")
	mux.Handle(grpcreflect.NewHandlerV1(reflector, connect.WithInterceptors(c.logging.LoggingInterceptor(), c.validation.ValidationInterceptor())))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector, connect.WithInterceptors(c.logging.LoggingInterceptor(), c.validation.ValidationInterceptor())))
	mux.Handle(mainv1connect.NewAdminServiceHandler(c.admin, connect.WithInterceptors(c.logging.LoggingInterceptor(), c.validation.ValidationInterceptor(), c.auth.AuthAdminServiceInterceptor())))
	mux.Handle(mainv1connect.NewProviderServiceHandler(c.provider, connect.WithInterceptors(c.logging.LoggingInterceptor(), c.validation.ValidationInterceptor(), c.auth.AuthProviderServiceInterceptor())))
	mux.Handle(mainv1connect.NewExternalInformationServiceHandler(c.extinfo, connect.WithInterceptors(c.logging.LoggingInterceptor(), c.validation.ValidationInterceptor(), c.auth.AuthExternalInformationServiceInterceptor())))
	mux.Handle(mainv1connect.NewKoyoServiceHandler(c.koyo, connect.WithInterceptors(c.logging.LoggingInterceptor(), c.validation.ValidationInterceptor(), c.auth.AuthKoyoServiceInterceptor())))
	mux.Handle(mainv1connect.NewHealthServiceHandler(c.health, connect.WithInterceptors(c.logging.LoggingInterceptor(), c.validation.ValidationInterceptor())))

	return mux
}

func (c *BeLifelineControllerImpl) Serve() error {
	err := c.cfg.LoadConfig()
	if err != nil {
		return err
	}

	mux := c.makeMux()
	servAddr := fmt.Sprintf("%s:%d", c.cfg.GetConfig().ListenAddr, c.cfg.GetConfig().Port)
	fmt.Printf("Listening on %s\n", servAddr)
	err = http.ListenAndServe(
		servAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

	return err
}

func (c *BeLifelineControllerImpl) TestServe() (*httptest.Server, error) {
	err := c.cfg.LoadConfig()
	if err != nil {
		return nil, err
	}

	mux := c.makeMux()
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	return server, nil
}
