package middleware

import (
	"fmt"
	"os"
	"time"

	"guiso/packages/logger"

	"github.com/labstack/echo/v4"
	"github.com/mattn/go-isatty"
)

func Logger() echo.MiddlewareFunc {
	log := logger.GetLogger()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			latency := time.Since(start)

			if logger.IsDev && isatty.IsTerminal(os.Stdout.Fd()) {
				// Development logging with colors
				statusColor := logger.ColorGreen
				if res.Status >= 300 && res.Status < 400 {
					statusColor = logger.ColorYellow
				} else if res.Status >= 400 {
					statusColor = logger.ColorRed
				}

				fmt.Printf("%s%s%s | %s%3d%s | %13v | %15s | %s%s%s %s\n",
					logger.ColorBlue, req.Method, logger.ColorReset,
					statusColor, res.Status, logger.ColorReset,
					latency,
					c.RealIP(),
					logger.ColorLavender, req.URL.Path, logger.ColorReset,
					req.URL.RawQuery,
				)
			} else {
				// Production logging or non-terminal output
				log.Info().
					Str("method", req.Method).
					Int("status", res.Status).
					Dur("latency", latency).
					Str("ip", c.RealIP()).
					Str("path", req.URL.Path).
					Str("query", req.URL.RawQuery).
					Msg("request")
			}

			return nil
		}
	}
}