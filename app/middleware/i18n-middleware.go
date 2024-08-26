package middleware

import (
	"guiso/app/utils"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func I18n() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bundle := i18n.NewBundle(language.English)
			bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
			
			loadMessageFiles(bundle, "locales", ".toml")

			lang := c.Param("lang")
			if lang == "" {
				lang = c.Request().Header.Get("Accept-Language")
			}

			localizer := i18n.NewLocalizer(bundle, lang)
			c.Set("localizer", localizer)

			markdownContent := loadMarkdownFiles("locales")
			c.Set("markdownContent", markdownContent)


			jsonLDScript, err := utils.GenerateJSONLD(localizer)
			if err != nil {
				// set default value
				jsonLDScript = "{}"
			}
			c.Set("jsonLDScript", jsonLDScript)

			return next(c)
		}
	}
}

func loadMessageFiles(bundle *i18n.Bundle, dir, ext string) {
	files, _ := filepath.Glob(filepath.Join(dir, "*"+ext))
	for _, file := range files {
		bundle.MustLoadMessageFile(file)
	}
}

func loadMarkdownFiles(dir string) map[string]map[string]templ.Component {
	files, _ := filepath.Glob(filepath.Join(dir, "*.*.md"))
	markdownContent := make(map[string]map[string]templ.Component)

	for _, file := range files {
		content, _ := ioutil.ReadFile(file)
		fileName := filepath.Base(file)
		parts := strings.Split(fileName, ".")
		if len(parts) != 3 {
			continue // Skip files not matching the pattern
		}
		key := parts[0]
		lang := parts[1]

		if markdownContent[key] == nil {
			markdownContent[key] = make(map[string]templ.Component)
		}
		markdownContent[key][lang] = utils.MarkdownToHTML(string(content))
	}

	return markdownContent
}