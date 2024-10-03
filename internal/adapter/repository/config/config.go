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
}

func NewConfigRepository() ConfigRepository {
	c := &configRepositoryImpl{}
	c.LoadConfig()
	return c
}

func (c *configRepositoryImpl) LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	err = envconfig.Process("kizuna", &c.config)
	return err
}

func (c *configRepositoryImpl) GetConfig() Config {
	return c.config
}
