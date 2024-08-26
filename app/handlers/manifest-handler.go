package handlers

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func LocalizedManifestHandler(c echo.Context) error {
	lang := c.Param("lang")
	filename := fmt.Sprintf("public/site.webmanifest.%s", lang)
	
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// Fall back to default manifest if localized version doesn't exist
		filename = "public/site.webmanifest"
	}
	
	return c.Attachment(filename, fmt.Sprintf("site.webmanifest.%s", lang))
}