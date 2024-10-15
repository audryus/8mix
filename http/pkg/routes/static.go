package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Static(app *fiber.App) {
	static := fiber.Static{
		Compress:      true,
		MaxAge:        86400,
		CacheDuration: time.Hour,
	}

	app.Static("/public/favicon", "./views/public/favicon", static)
	app.Static("/public/css", "./views/public/css", static)
	app.Static("/public/fonts", "./views/public/fonts", static)
	app.Static("/public/images", "./views/public/images", static)
	app.Static("/public/js", "./views/public/js", static)
	app.Static("/public/webfonts", "./views/public/webfonts", static)
}
