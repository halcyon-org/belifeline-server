package ent

import (
	"fmt"

	"github.com/halcyon-org/kizuna/ent"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
)

func InitDB(cfg config.ConfigRepository) (*ent.Client, error) {
	config := cfg.GetConfig()

	client, err := ent.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Postgres.User, config.Postgres.Password, config.Postgres.Host, config.Postgres.Port, config.Postgres.DB))
	if err != nil {
		return nil, err
	}

	defer client.Close()
	return client, nil
}
