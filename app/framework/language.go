package framework

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/language"
)

var supportedLanguages = []language.Tag{
	language.Spanish,
	language.English,
	// Add more supported languages here
}

func LanguageMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		lang := c.Param("lang")
		if lang == "" {
			// Detect language from Accept-Language header
			accept := c.Request().Header.Get("Accept-Language")
			tag, _ := language.MatchStrings(language.NewMatcher(supportedLanguages), accept)
			lang = tag.String()
		}

		// Validate language
		tag, err := language.Parse(lang)
		if err != nil || !isSupported(tag) {
			// Fallback to default language (Español)
			lang = "es"
		} else {
			lang = tag.String()
		}

		// Set language in context
		c.Set("lang", lang)

		// Set language cookie
		cookie := new(http.Cookie)
		cookie.Name = "lang"
		cookie.Value = lang
		cookie.Path = "/"
		c.SetCookie(cookie)

		return next(c)
	}
}

func isSupported(tag language.Tag) bool {
	for _, supported := range supportedLanguages {
		if tag == supported {
			return true
		}
	}
	return false
}