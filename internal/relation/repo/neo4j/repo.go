package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.opentelemetry.io/otel/trace"
)

type repo struct {
	cfg    Config
	driver neo4j.DriverWithContext
}

func New(cfg Config, driver neo4j.DriverWithContext, tr trace.Tracer) tracedRepo {
	return tracedRepo{
		repo: repo{
			cfg:    cfg,
			driver: driver,
		},
		tracer: tr,
	}
}
