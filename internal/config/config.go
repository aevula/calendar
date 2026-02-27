package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	Dev  = "dev"
	Test = "test"
	CI   = "ci"
	Prod = "prod"
)

type Config struct {
	Env string `yaml:"env" env:"ENV"`

	Server   Server   `yaml:"server"`
	Backjobs Backjobs `yaml:"worker"`
	Consumer Consumer `yaml:"consumer"`
	Queues   Queues   `yaml:"queues"`

	Log Log `yaml:"log"`
	DB  DB  `yaml:"db"`
}

func (cfg Config) IsDev() bool {
	return cfg.Env == Dev
}

func (cfg Config) IsTest() bool {
	return cfg.Env == Test
}

func (cfg Config) IsCI() bool {
	return cfg.Env == CI
}

func (cfg Config) IsProd() bool {
	return cfg.Env == Prod
}

func MustLoad() Config {
	cfg := Config{}

	cfgPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		log.Fatalf("environment variable CONFIG_PATH must be set")
	}

	if _, err := os.Stat(cfgPath); err != nil {
		log.Fatalf("error loading config file %v", err)
	}

	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		log.Fatalf("error loading config %v", err)
	}

	return cfg
}
