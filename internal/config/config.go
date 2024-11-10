package config

import (
	grpc "github.com/Karzoug/meower-relation-service/internal/delivery/grpc/server"
	httpConfig "github.com/Karzoug/meower-relation-service/internal/delivery/http/config"
	relRepo "github.com/Karzoug/meower-relation-service/internal/relation/repo/neo4j"
	"github.com/Karzoug/meower-relation-service/pkg/metric/prom"
	"github.com/Karzoug/meower-relation-service/pkg/neo4j"
	"github.com/Karzoug/meower-relation-service/pkg/trace/otlp"

	"github.com/rs/zerolog"
)

type Config struct {
	LogLevel     zerolog.Level           `env:"LOG_LEVEL" envDefault:"info"`
	HTTP         httpConfig.ServerConfig `envPrefix:"HTTP_"`
	GRPC         grpc.Config             `envPrefix:"GRPC_"`
	PromHTTP     prom.ServerConfig       `envPrefix:"PROM_"`
	OTLP         otlp.Config             `envPrefix:"OTLP_"`
	Neo4j        neo4j.Config            `envPrefix:"NEO4J_"`
	RelationRepo relRepo.Config          `envPrefix:"RELATION_REPO_"`
}
