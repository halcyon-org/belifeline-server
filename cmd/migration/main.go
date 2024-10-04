package main

import (
	"context"
	"fmt"
	"log"

	"github.com/halcyon-org/kizuna/ent"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.NewConfigRepository()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	config := cfg.GetConfig()

	client, err := ent.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Postgres.User, config.Postgres.Password, config.Postgres.Host, config.Postgres.Port, config.Postgres.DB))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
