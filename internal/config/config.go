package config

import (
	httpConfig "github.com/Karzoug/meower-relation-service/internal/delivery/http/config"
	"github.com/Karzoug/meower-relation-service/pkg/trace/otlp"

	"github.com/rs/zerolog"
)

type Config struct {
	LogLevel zerolog.Level           `env:"LOG_LEVEL" envDefault:"info"`
	HTTP     httpConfig.ServerConfig `envPrefix:"HTTP_"`
	OTLP     otlp.Config             `envPrefix:"OTLP_"`
}
