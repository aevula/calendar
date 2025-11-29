package config

import "time"

type Server struct {
	Port int `yaml:"port" env:"SERVER_PORT"`

	IdleTimeout  time.Duration `yaml:"idle_timeout"  env:"SERVER_IDLE_TIMEOUT"`
	ReadTimeout  time.Duration `yaml:"read_timeout"  env:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"SERVER_WRITE_TIMEOUT"`

	InitTimeout time.Duration `yaml:"init_timeout" env:"SERVER_INIT_TIMEOUT"`
	ShutTimeout time.Duration `yaml:"shut_timeout" env:"SERVER_SHUT_TIMEOUT"`
}
