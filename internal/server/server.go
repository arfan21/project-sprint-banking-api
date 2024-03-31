package server

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/arfan21/project-sprint-banking-api/config"
	_ "github.com/arfan21/project-sprint-banking-api/docs"
	"github.com/arfan21/project-sprint-banking-api/pkg/exception"
	"github.com/arfan21/project-sprint-banking-api/pkg/logger"
	"github.com/arfan21/project-sprint-banking-api/pkg/middleware"
	"github.com/arfan21/project-sprint-banking-api/pkg/pkgutil"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog"
)

const (
	ctxTimeout = 5
)

type Server struct {
	app *fiber.App
	db  *sql.DB
}

func New(
	db *sql.DB,
) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.FiberErrorHandler,
	})

	// app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	timeout := time.Duration(config.Get().Service.Timeout) * time.Second
	app.Use(middleware.Timeout(timeout))

	prom := fiberprometheus.New(config.Get().Service.Name)
	prom.RegisterAt(app, "/metrics")
	app.Use(prom.Middleware)

	app.Use(cors.New())

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		// Logger: logger.Log(context.Background()),
		GetLogger: func(c *fiber.Ctx) zerolog.Logger {
			return *logger.Log(c.UserContext())
		},
	}))

	app.Use(recover.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	return &Server{
		app: app,
		db:  db,
	}
}

func (s *Server) Run() error {
	s.Routes()
	s.app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(pkgutil.HTTPResponse{
			Message: "Not Found",
		})
	})
	ctx := context.Background()
	go func() {
		if err := s.app.Listen(pkgutil.GetPort()); err != nil {
			logger.Log(ctx).Fatal().Err(err).Msg("failed to start server")
		}
	}()

	// go func() {
	// 	logger.Log(ctx).Info().Msgf("Starting prometheus exporter on port %s", config.Get().Otel.ExporterPrometheusPort)
	// 	http.Handle(config.Get().Otel.ExporterPrometheusPath)
	// 	if err := http.ListenAndServe(pkgutil.GetPort(config.Get().Otel.ExporterPrometheusPort), nil); err != nil {
	// 		logger.Log(ctx).Fatal().Err(err).Msg("failed to start prometheus exporter")
	// 	}
	// }()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	_, shutdown := context.WithTimeout(ctx, ctxTimeout*time.Second)
	defer shutdown()

	logger.Log(ctx).Info().Msg("shutting down server")
	return s.app.Shutdown()
}
