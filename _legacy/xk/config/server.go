package config

import (
	iconfig "github.com/dyweb/go.ice/config"
)

type ServerConfig struct {
	Http iconfig.HttpServerConfig `yaml:"http"`
	Grpc iconfig.GrpcServerConfig `yaml:"grpc"`
}
