package routes

import (
	"github.com/labstack/echo/v4"
)

type ApiResponse struct {
	Message string `json:"message"`
	Code    int
}

func Api(e *echo.Echo) {
	e.GET("/api", func(c echo.Context) error {
		return c.JSON(200, ApiResponse{
			Message: "Hello World",
			Code:    200,
		})
	})
}
