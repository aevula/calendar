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
	Env string `yaml:"env" env:"ENV"`

	Server Server `yaml:"server"`
	Log    Log    `yaml:"log"`
	Db     Db     `yaml:"db"`
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

type Server struct {
	Port int `yaml:"port"          env:"SERVER_PORT"`

	IdleTimeout  time.Duration `yaml:"idle_timeout"  env:"SERVER_IDLE_TIMEOUT"`
	ReadTimeout  time.Duration `yaml:"read_timeout"  env:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"SERVER_WRITE_TIMEOUT"`

	InitTimeout time.Duration `yaml:"init_timeout" env:"SERVER_INIT_TIMEOUT"`
	ShutTimeout time.Duration `yaml:"shut_timeout" env:"SERVER_SHUT_TIMEOUT"`
}

type Log struct {
	Level string `yaml:"level" env:"LOG_LEVEL"`
	Trace bool   `yaml:"trace" env:"LOG_TRACE"`
	Plain bool   `yaml:"plain" env:"LOG_PLAIN"`
}

type Db struct {
	User     string `yaml:"user"     env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	Host     string `yaml:"host"     env:"DB_HOST"`
	Port     int    `yaml:"port"     env:"DB_PORT"`
	Name     string `yaml:"name"     env:"DB_NAME"`
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
