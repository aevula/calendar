package config

import "time"

type Consumer struct {
	Seeds  []string `yaml:"seeds"`
	Topics []string `yaml:"topics"`

	InitTimeout time.Duration `yaml:"init_timeout" env:"CONSUMER_INIT_TIMEOUT"`
	ShutTimeout time.Duration `yaml:"shut_timeout" env:"CONSUMER_SHUT_TIMEOUT"`
}
