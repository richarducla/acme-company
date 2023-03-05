package config

import (
	"fmt"

	"github.com/joeshaw/envdecode"
)

type Config struct {
	SrvHost    string `env:"HOST,default=0.0.0.0"`
	SrvPort    string `env:"PORT,default=8080"`
	Env        string `env:"APP_ENV,default=dev"`
	DbHost     string `env:"DB_HOST,default=localhost"`
	DbPort     string `env:"DB_PORT,default=5432"`
	DbUser     string `env:"DB_USER,default=root"`
	DbPassword string `env:"DB_PASSWORD,default=secret"`
	DbName     string `env:"DB_NAME,default=acme"`
}

func FromEnv() (Config, error) {
	c := Config{}
	err := envdecode.Decode(&c)
	return c, err
}

func (c Config) SrvAddr() string {
	return fmt.Sprintf("%s:%s", c.SrvHost, c.SrvPort)
}
