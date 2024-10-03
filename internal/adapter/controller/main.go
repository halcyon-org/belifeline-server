package controller

import (
	"fmt"
	"net/http"

	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type BeLifelineController interface {
	Serve() error
}

type BeLifelineControllerImpl struct {
	server mainv1connect.BeLifelineServiceHandler

	cfg config.ConfigRepository
}

func NewBeLifelineController(server mainv1connect.BeLifelineServiceHandler, config config.ConfigRepository) BeLifelineController {
	return &BeLifelineControllerImpl{
		server: server,
		cfg:    config,
	}
}

func (c *BeLifelineControllerImpl) Serve() error {
	err := c.cfg.LoadConfig()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	path, handler := mainv1connect.NewBeLifelineServiceHandler(c.server)
	mux.Handle(path, handler)

	err = http.ListenAndServe(
		fmt.Sprintf("%s:%d", c.cfg.GetConfig().ListenAddr, c.cfg.GetConfig().Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

	return err
}
