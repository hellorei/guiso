package middleware

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)



func Static() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path
	
			// Check if the request is for a static file
			if	strings.HasPrefix(path, "/css/") || 
					strings.HasPrefix(path, "/js/") || 
					strings.HasPrefix(path, "/out/") || 
					strings.HasPrefix(path, "/images/") || 
					strings.HasPrefix(path, "/favicon/") || 
					strings.HasPrefix(path, "/fonts/") {
				file := filepath.Join("public", path)
				if _, err := os.Stat(file); err == nil {
					return c.File(file)
				}
			}
	
			return next(c)
		}
	}
}