package config

import (
	iconfig "github.com/at15/go.ice/ice/config"
)

type ServerConfig struct {
	Http iconfig.HttpServerConfig `yaml:"http"`
	Grpc iconfig.GrpcServerConfig `yaml:"grpc"`
}
