package config

type Log struct {
	Level string `yaml:"level" env:"LOG_LEVEL"`
	Trace bool   `yaml:"trace" env:"LOG_TRACE"`
	Plain bool   `yaml:"plain" env:"LOG_PLAIN"`
}
