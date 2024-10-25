package ent

import (
	"net/url"

	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
	_ "github.com/lib/pq"
)

func InitDB(cfg config.Repository) (*ent.Client, error) {
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
		return nil, err
	}

	return client, nil
}
