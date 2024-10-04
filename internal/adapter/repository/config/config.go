package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type ConfigRepository interface {
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

func NewConfigRepository() (ConfigRepository, error) {
	c := &configRepositoryImpl{}
	err := c.LoadConfig()
	return c, err
}

func (c *configRepositoryImpl) LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	psqlConfig := PostgresConfig{}
	err = envconfig.Process("postgres", &psqlConfig)
	if err != nil {
		return err
	}

	err = envconfig.Process("kizuna", &c.config)
	if err != nil {
		return err
	}

	c.config.Postgres = psqlConfig
	return nil
}

func (c *configRepositoryImpl) GetConfig() Config {
	return c.config
}
