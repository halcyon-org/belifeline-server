package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"github.com/halcyon-org/kizuna/internal/adapter/api"
	"github.com/halcyon-org/kizuna/internal/infrastructure/repository/config"
)

func main() {
	server := &api.BeLifelineServer{}
	cfg := config.NewConfigRepository()
	err := cfg.LoadConfig()
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	path, handler := mainv1connect.NewBeLifelineServiceHandler(server)
	mux.Handle(path, handler)
	http.ListenAndServe(
		fmt.Sprintf("%s:%d", cfg.GetConfig().ListenAddr, cfg.GetConfig().Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
