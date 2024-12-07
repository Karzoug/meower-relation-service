package config

import (
	"github.com/Karzoug/meower-common-go/metric/prom"
	"github.com/Karzoug/meower-common-go/trace/otlp"

	grpc "github.com/Karzoug/meower-relation-service/internal/delivery/grpc/server"
	relRepo "github.com/Karzoug/meower-relation-service/internal/relation/repo/neo4j"
	"github.com/Karzoug/meower-relation-service/pkg/neo4j"

	"github.com/rs/zerolog"
)

type Config struct {
	LogLevel     zerolog.Level     `env:"LOG_LEVEL" envDefault:"info"`
	GRPC         grpc.Config       `envPrefix:"GRPC_"`
	PromHTTP     prom.ServerConfig `envPrefix:"PROM_"`
	OTLP         otlp.Config       `envPrefix:"OTLP_"`
	Neo4j        neo4j.Config      `envPrefix:"NEO4J_"`
	RelationRepo relRepo.Config    `envPrefix:"RELATION_REPO_"`
}
