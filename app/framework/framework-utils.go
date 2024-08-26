package framework

import (
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CachedComponent struct {
	Component    templ.Component
	RevalidateAt int64 // Unix timestamp
}

var cache = map[string]CachedComponent{}

// EchoRouter is an interface that both echo.Echo and echo.Group satisfy
type EchoRouter interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	// Add other HTTP methods as needed
}

func RegisterComponent(e EchoRouter, url string, renderMethod func(ctx echo.Context) templ.Component, revalidate int64, method string) {
	handler := func(ctx echo.Context) error {
		startTime := time.Now()
		path := ctx.Request().URL.Path
		var template templ.Component
		shouldRender := true
		timestamp := int64(time.Now().Unix())
		if cachedTemplate, ok := cache[path]; ok {
			template = cachedTemplate.Component
			if revalidate == -1 || timestamp < cachedTemplate.RevalidateAt {
				shouldRender = false
			}
		}

		// revalidation period has passed or not cached yet, render the component
		if shouldRender {
			println("[rendering-static-path]: " + path)
			renderedComponent := renderMethod(ctx)
			if renderedComponent != nil {
				cache[path] = CachedComponent{
					Component:    renderedComponent,
					RevalidateAt: timestamp + revalidate,
				}
				template = cache[path].Component
			} else {
				return ctx.String(500, "Render method did not return a component")
			}
		}

		err := template.Render(ctx.Request().Context(), ctx.Response().Writer)
		if err != nil {
			return ctx.String(500, "Error rendering component: "+err.Error())
		}
		println("[rendered]: " + path + " in " + time.Since(startTime).String())
		return nil
	}

	switch method {
	case "GET":
		e.GET(url, handler)
	case "POST":
		e.POST(url, handler)
	// Add other HTTP methods as needed
	default:
		panic("Unsupported HTTP method: " + method)
	}
}

func RenderView(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}