package middleware

import (
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func trackPageView(c echo.Context) {
	ga4MeasurementID := os.Getenv("GA4_MEASUREMENT_ID")
	ga4APISecret := os.Getenv("GA4_API_SECRET")

	clientIP := c.RealIP()
	userAgent := c.Request().UserAgent()
	path := c.Request().URL.Path

	data := url.Values{}
	data.Set("measurement_id", ga4MeasurementID)
	data.Set("api_secret", ga4APISecret)
	data.Set("client_id", clientIP)
	data.Set("user_agent", userAgent)
	data.Set("page_location", path)
	data.Set("page_title", path)

	go http.PostForm("https://www.google-analytics.com/mp/collect", data)
}

func GoogleAnalytics() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !strings.HasPrefix(c.Path(), "/static") && !strings.HasPrefix(c.Path(), "/favicon") {
				trackPageView(c)
			}
			return next(c)
		}
	}
}