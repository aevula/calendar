package config

import "time"

type Backjobs struct {
	WorkersCount int           `yaml:"workers_count" env:"BACKJOBS_WORKERS_COUNT"`
	RefreshRate  time.Duration `yaml:"refresh_rate"  env:"BACKJOBS_REFRESH_RATE"`

	InitTimeout time.Duration `yaml:"init_timeout" env:"BACKJOBS_INIT_TIMEOUT"`
	ShutTimeout time.Duration `yaml:"shut_timeout" env:"BACKJOBS_SHUT_TIMEOUT"`
}
