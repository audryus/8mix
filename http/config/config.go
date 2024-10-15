package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Server `yaml:"server"`
		App    `yaml:"app"`
		Http   `yaml:"http"`
	}

	Server struct {
		Header string `yaml:"header"`
		Addr   string `env-required:"true" yaml:"addr" env:"8MIX_SERVER_ADDR"`
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	Http struct {
		Addr string `env-required:"true" yaml:"addr" env:"8MIX_HTTP_ADDR"`
	}
)

func New() Config {
	cfg := Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", &cfg)
	if err != nil {
		log.Fatalf("config load error: %s", err)
		panic(err)
	}

	return cfg
}
