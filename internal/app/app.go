package app

import (
	"context"
	"runtime"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"golang.org/x/sync/errgroup"

	"github.com/Karzoug/meower-relation-service/internal/config"
	healthHandlers "github.com/Karzoug/meower-relation-service/internal/delivery/http/handler/health"
	relHandlers "github.com/Karzoug/meower-relation-service/internal/delivery/http/handler/relation"
	httpServer "github.com/Karzoug/meower-relation-service/internal/delivery/http/server"
	"github.com/Karzoug/meower-relation-service/internal/relation/service"
	"github.com/Karzoug/meower-relation-service/pkg/buildinfo"
	"github.com/Karzoug/meower-relation-service/pkg/healthcheck"
	"github.com/Karzoug/meower-relation-service/pkg/trace/otlp"
)

const (
	serviceName     = "RelationService"
	metricNamespace = "relation_service"
	pkgName         = "github.com/Karzoug/meower-relation-service"
	initTimeout     = 10 * time.Second
)

var serviceVersion = buildinfo.Get().ServiceVersion

func Run(ctx context.Context, logger zerolog.Logger) error {
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(cfg.LogLevel)

	logger.Info().
		Int("GOMAXPROCS", runtime.GOMAXPROCS(0)).
		Str("log level", cfg.LogLevel.String()).
		Msg("starting up")

	// set timeout for initialization
	ctxInit, closeCtx := context.WithTimeout(ctx, initTimeout)
	defer closeCtx()

	// set up tracer
	cfg.OTLP.ServiceName = serviceName
	cfg.OTLP.ServiceVersion = serviceVersion
	cfg.OTLP.ExcludedRoutes = map[string]struct{}{
		"/readiness": {},
		"/liveness":  {},
	}
	shutdownTracer, err := otlp.RegisterGlobal(ctxInit, cfg.OTLP)
	if err != nil {
		return err
	}
	defer doClose(shutdownTracer, logger)

	tracer := otel.GetTracerProvider().Tracer(pkgName)

	// set up service
	relationService := service.NewRelationService()

	// set up http server
	httpSrv := httpServer.New(
		cfg.HTTP,
		[]httpServer.Routes{
			relHandlers.RoutesFunc(relationService, tracer, logger),
			healthHandlers.RoutesFunc(healthcheck.New(), logger),
		},
		logger)

	eg, ctx := errgroup.WithContext(ctx)
	// run service http server
	eg.Go(func() error {
		return httpSrv.Run(ctx)
	})

	return eg.Wait()
}

func doClose(fn func(context.Context) error, logger zerolog.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := fn(ctx); err != nil {
		logger.Error().
			Err(err).
			Msg("error closing")
	}
}
