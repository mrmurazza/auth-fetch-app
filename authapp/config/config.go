package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	DBHost     string `envconfig:"DB_HOST" default:"authapp.db"`
	DBDriver   string `envconfig:"DB_DRIVER" default:"sqlite3"`
	DBUser     string `envconfig:"DB_USER" default:""`
	DBPassword string `envconfig:"DB_PASSWORD" default:""`
	DBPort     string `envconfig:"DB_PORT" default:""`
	DBName     string `envconfig:"DB_NAME" default:"auth-fetch-app"`
}

func Get() *Config {
	c := Config{}
	once.Do(func() {
		envconfig.MustProcess("", &c)
		cfg = &c
	})

	return cfg
}
