package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	Dev  = "dev"
	Test = "test"
	CI   = "ci"
	Prod = "prod"
)

type Config struct {
	Env    string `yaml:"env"    env:"ENV"`
	Server Server `yaml:"server"`
	Log    Log    `yaml:"log"`
}

type Server struct {
	Port         int           `yaml:"port"          env:"SERVER_PORT"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"  env:"SERVER_IDLE_TIMEOUT"`
	ReadTimeout  time.Duration `yaml:"read_timeout"  env:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"SERVER_WRITE_TIMEOUT"`
}

type Log struct {
	Level string `yaml:"level" env:"LOG_LEVEL"`
	Trace bool   `yaml:"trace" env:"LOG_TRACE"`
	Plain bool   `yaml:"plain" env:"LOG_PLAIN"`
}

func MustLoad() Config {
	cfg := Config{}

	cfgPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		log.Fatalf("environment variable CONFIG_PATH must be set")
	}

	if _, err := os.Stat(cfgPath); err != nil {
		log.Fatalf("error loading config %v", err)
	}

	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		log.Fatalf("error loading config %v", err)
	}

	return cfg
}
