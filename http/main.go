package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/audryus/8mix/http/config"
	"github.com/audryus/8mix/http/internal/controller"
	"github.com/audryus/8mix/http/internal/domain/playlist"
	"github.com/audryus/8mix/http/internal/domain/track"
	"github.com/audryus/8mix/http/internal/usecase"
	"github.com/audryus/8mix/http/pkg/logger"
	"github.com/audryus/8mix/http/pkg/middleware"
	"github.com/audryus/8mix/http/pkg/mongo"
	"github.com/audryus/8mix/http/pkg/routes"
	"github.com/audryus/8mix/http/pkg/temporal"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(CreateServer),
		Register(),
		Routes(),
		fx.Invoke(func(lifecycle fx.Lifecycle, logger *logger.Log, app *fiber.App, cfg config.Config, mongo *mongo.Mongo) {
			lifecycle.Append(fx.Hook{
				OnStart: func(context.Context) error {
					go app.Listen(cfg.Http.Addr)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					logger.Info("shuting down")
					mongo.Close()
					return app.Shutdown()
				},
			})
		}),
	)

	app.Run()
}

// CreateServer creates a new GoFiber server instance
func CreateServer(cfg config.Config) *fiber.App {
	engine := django.NewFileSystem(http.Dir("./views"), ".html")

	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ServerHeader: cfg.Server.Header,
		AppName:      cfg.App.Name,
		Views:        engine,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	return app
}

func Routes() fx.Option {
	return fx.Options(
		fx.Invoke(middleware.Fiber),
		fx.Invoke(routes.Static),
		fx.Invoke(controller.New),
	)
}

func Register() fx.Option {
	return fx.Options(
		fx.Provide(
			config.New,
			logger.New,
			temporal.New,
			mongo.New,
			fx.Annotate(track.NewTrackRepo, fx.As(new(track.ITrackRepo))),
			fx.Annotate(track.NewTrackService, fx.As(new(usecase.ITrackService))),
			fx.Annotate(playlist.NewPlaylistRepo, fx.As(new(playlist.IPlaylistRepo))),
			fx.Annotate(playlist.NewPlaylisteWorkflow, fx.As(new(playlist.IPlaylistWorkflow))),
			fx.Annotate(playlist.NewPlaylistService, fx.As(new(usecase.IPlaylistService))),
			usecase.NewPlaylistUC,
		),
	)
}

func checkIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
