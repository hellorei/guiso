package middleware

import (
	"guiso/app/types"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func SEO(defaultData types.SEOData) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			localizer := c.Get("localizer").(*i18n.Localizer)
			
			seoData := types.SEOData{
				Title:      	 	localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Meta.Title"}),
				Description:	 	localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Meta.Description"}),
				URL:       	  	c.Scheme() + "://" + c.Request().Host + c.Request().URL.Path,
				Image:       		c.Scheme() + "://" + c.Request().Host + "/images/og-image.jpg",
				Type:       	 	"website",
				Locale:    	  	c.Request().Header.Get("Accept-Language"),
				SiteName:    		localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Meta.Title"}),
				TwitterHandle: 	"@guiso",
			}

			c.Set("seo", seoData)
			return next(c)
		}
	}
}