package neo4j

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/auth"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/config"
	"github.com/rs/zerolog"
)

func NewDriver(ctx context.Context, cfg Config, zlog zerolog.Logger) (neo4j.DriverWithContext, error) {
	var auth auth.TokenManager
	if cfg.Username == "" {
		auth = neo4j.NoAuth()
	} else {
		auth = neo4j.BasicAuth(cfg.Username, cfg.Password, "")
	}

	zlog = zlog.With().Str("component", "neo4j driver").Logger()
	driver, err := neo4j.NewDriverWithContext(cfg.URI, auth, func(c *config.Config) {
		c.Log = logger{l: zlog}
	})
	if err != nil {
		return driver, fmt.Errorf("failed to create neo4j driver: %w", err)
	}

	if err := driver.VerifyConnectivity(ctx); err != nil {
		return nil, fmt.Errorf("failed to verify neo4j connectivity: %w", err)
	}

	return driver, nil
}
