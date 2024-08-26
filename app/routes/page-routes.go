package routes

import (
	"guiso/app/framework"
	"guiso/app/templ/components"
	"guiso/app/templ/pages"
	"guiso/app/types"
	"guiso/app/utils"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var Globals = map[string]interface{}{
	"Clicked": 0,
}

func createDefaultContext(c echo.Context) types.AppContext {
	path := c.Request().URL.Path
	seo := c.Get("seo").(types.SEOData)
	seo.URL = c.Scheme() + "://" + c.Request().Host + path
	
	currentLocale := utils.GetCurrentLocale(path)
	localizer := c.Get("localizer").(*i18n.Localizer)
	markdownContent, _ := c.Get("markdownContent").(map[string]map[string]templ.Component)


	jsonLD, err := utils.GenerateJSONLD(localizer)
	if err != nil {
		// Handle error (maybe log it)
		jsonLD = "{}" // or some default value
	}

	return types.AppContext{
		Localizer:       localizer,
		CurrentPath:     path,
		CurrentLocale:   currentLocale,
		Lang:            currentLocale,
		SEO:             seo,
		MarkdownContent: markdownContent,
		JSONLDScript: jsonLD,
	}
}

func Pages(e *echo.Echo) {
	// Root redirect
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(302, "/es")
	})

	// Language-specific routes
	langGroup := e.Group("/:lang")
	langGroup.Use(framework.LanguageMiddleware)

	framework.RegisterComponent(langGroup, "", func(c echo.Context) templ.Component {
		appCtx := createDefaultContext(c)
    appCtx.SEO = utils.UpdateSEO(appCtx.SEO, map[string]string{
        "title": appCtx.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Meta.Title"}) + "| Home",
        "description": appCtx.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Meta.Description"}),
    })
		return pages.Index(appCtx)
	}, 60, "GET")

	framework.RegisterComponent(langGroup, "/pokemon", func(c echo.Context) templ.Component {
		var result types.PokemonList
		err := utils.Get("https://pokeapi.co/api/v2/pokemon", &result)

		if err != nil {
			return nil
		}

		appCtx := createDefaultContext(c)
		appCtx.Meta.Title = "Pokemon List"
		appCtx.Meta.Description = "Explore our Pokemon list"
		appCtx.Meta.Keywords = "pokemon, list, explore"
		return pages.Pokemon(appCtx, result.Results)
	}, -1, "GET")

	framework.RegisterComponent(langGroup, "/pokemon/:name", func(c echo.Context) templ.Component {
		var result types.Pokemon
		err := utils.Get("https://pokeapi.co/api/v2/pokemon/"+c.Param("name"), &result)

		if err != nil {
			return nil
		}
		
		appCtx := createDefaultContext(c)
		appCtx.Meta.Title = "Pokemon Detail"
		appCtx.Meta.Description = "Explore our Pokemon detail"
		appCtx.Meta.Keywords = "pokemon, detail, explore"
		return pages.PokemonSlug(appCtx, result)
	}, -1, "GET")

	framework.RegisterComponent(langGroup, "/clicked", func(c echo.Context) templ.Component {
		Globals["Clicked"] = Globals["Clicked"].(int) + 1
		
		appCtx := createDefaultContext(c)
		appCtx.Meta.Title = "Clicked"
		appCtx.Meta.Description = "Clicked"
		appCtx.Meta.Keywords = "clicked"
		return components.Clicked(appCtx, Globals["Clicked"].(int))
	}, 0, "POST")
}