package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Repository interface {
	LoadConfig() error
	GetConfig() Config
}

type configRepositoryImpl struct {
	config Config
}

type Config struct {
	Port       int
	ListenAddr string `split_words:"true"`

	Postgres PostgresConfig
}

type PostgresConfig struct {
	DB       string
	Host     string
	User     string
	Password string
	Port     int
}

func NewConfigRepository() (Repository, error) {
	c := &configRepositoryImpl{}
	err := c.LoadConfig()

	return c, err
}

func (c *configRepositoryImpl) LoadConfig() error {
	env := os.Getenv("ENV")
	if env != "production" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	psqlConfig := PostgresConfig{}
	if err := envconfig.Process("postgres", &psqlConfig); err != nil {
		return err
	}

	if err := envconfig.Process("kizuna", &c.config); err != nil {
		return err
	}

	c.config.Postgres = psqlConfig

	return nil
}

func (c *configRepositoryImpl) GetConfig() Config {
	return c.config
}
