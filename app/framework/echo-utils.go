package framework

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"guiso/app/handlers"
	guiso_middleware "guiso/app/middleware"
	"guiso/app/utils"
)

func SetupEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(guiso_middleware.Static())
	e.Use(guiso_middleware.Logger())
	e.Use(guiso_middleware.GoogleAnalytics())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(guiso_middleware.I18n())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		localizer := c.Get("localizer").(*i18n.Localizer)
		baseURL := c.Scheme() + "://" + c.Request().Host
		c.Set("seo", utils.GetDefaultSEO(localizer, baseURL))
		return next(c)
	}
})

	e.Static("/favicon", "public/favicon")
	e.GET("/robots.txt", handlers.RobotsHandler)

	// Serve the default manifest
	e.GET("/site.webmanifest", func(c echo.Context) error {
		return c.Attachment("public/site.webmanifest", "site.webmanifest")
	})

	// Serve localized manifests
	e.GET("/:lang/site.webmanifest", handlers.LocalizedManifestHandler)

	e.GET("/ws", handlers.WebSocketHandler)

	return e
}

func StartServer(e *echo.Echo) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	e.Logger.Fatal(e.Start(":" + port))
	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}
