package main

import (
	"context"
	"flag"
	"log"
	"net/url"
	"os"

	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
	_ "github.com/lib/pq"
)

func main() {
	dry := flag.Bool("dry", false, "dry run")
	flag.Parse()

	cfg, err := config.NewConfigRepository()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	config := cfg.GetConfig()

	dataSource := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(config.Postgres.User, config.Postgres.Password),
		Host:     config.Postgres.Host,
		Path:     config.Postgres.DB,
		RawQuery: "sslmode=disable",
	}

	client, err := ent.Open(
		"postgres",
		dataSource.String(),
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if *dry {
		if err := client.Schema.WriteTo(context.Background(), os.Stdout); err != nil {
			log.Fatalf("failed printing schema changes: %v", err)
		}
	} else {
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}
}
