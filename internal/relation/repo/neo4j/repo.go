package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type repo struct {
	cfg    Config
	driver neo4j.DriverWithContext
}

func New(cfg Config, driver neo4j.DriverWithContext) repo {
	return repo{
		cfg:    cfg,
		driver: driver,
	}
}
