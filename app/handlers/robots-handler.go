package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func RobotsHandler(c echo.Context) error {
	host := c.Request().Host

	isStaging := strings.HasPrefix(host, "staging.") || strings.Contains(host, "-staging.")
	isDev := strings.HasPrefix(host, "dev.") || strings.Contains(host, "-dev") || strings.HasPrefix(host, "dev.") || strings.Contains(host, "-dev.")

	if isStaging || isDev {
		return c.String(http.StatusOK, "User-agent: *\nDisallow: /")
	}

	sitemapURL := fmt.Sprintf("https://%s/sitemap.xml", host)
	robotsTxt := fmt.Sprintf(`User-agent: *
Allow: /

Sitemap: %s`, sitemapURL)

	return c.String(http.StatusOK, robotsTxt)
}