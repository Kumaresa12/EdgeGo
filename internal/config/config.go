package config

import "time"

type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Logging LoggingConfig `yaml:"logging"`
}

type ServerConfig struct {
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	IdleTimeout  time.Duration `yaml:"idleTimeout"`
}

type LoggingConfig struct {
	level string `yaml:"level"`
}
