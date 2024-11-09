package neo4j

import (
	"fmt"

	"github.com/rs/zerolog"
)

type logger struct {
	l zerolog.Logger
}

func (l logger) Error(name string, id string, err error) {
	l.l.Error().
		Str("name", name).
		Str("id", id).
		Err(err).
		Msg("neo4j error")
}

func (l logger) Warnf(name string, id string, msg string, args ...any) {
	l.l.Warn().
		Str("name", name).
		Str("id", id).
		Str("args", fmt.Sprint(args...)).
		Msg("neo4j warning")
}

func (l logger) Infof(name string, id string, msg string, args ...any) {
}

func (l logger) Debugf(name string, id string, msg string, args ...any) {
}
