package middleware

import (
	"time"

	"github.com/audryus/8mix/http/pkg/logger"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Fiber(app *fiber.App, logger *logger.Log) {
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger.Core(),
	}))
	app.Use(compress.New())
	app.Use(cache.New(cache.Config{
		Expiration:   24 * time.Hour,
		CacheControl: true,
	}))

	app.Use(requestid.New(requestid.Config{
		Header: "x-kong-request-id",
	}))

}
