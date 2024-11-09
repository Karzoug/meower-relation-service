package config

import (
	grpc "github.com/Karzoug/meower-relation-service/internal/delivery/grpc/server"
	httpConfig "github.com/Karzoug/meower-relation-service/internal/delivery/http/config"
	"github.com/Karzoug/meower-relation-service/pkg/trace/otlp"

	"github.com/rs/zerolog"
)

type Config struct {
	LogLevel     zerolog.Level           `env:"LOG_LEVEL" envDefault:"info"`
	HTTP         httpConfig.ServerConfig `envPrefix:"HTTP_"`
	GRPC         grpc.Config             `envPrefix:"GRPC_"`
	OTLP         otlp.Config             `envPrefix:"OTLP_"`
	Neo4j        neo4j.Config            `envPrefix:"NEO4J_"`
}
