package types

import (
	"github.com/a-h/templ"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Meta struct {
	Title string
	Description string
	Keywords string
}

type SEOData struct {
	Title       	string
	Description 	string
	URL         	string
	Image       	string
	Type        	string
	Locale				string
	SiteName    	string
	TwitterHandle string
	FacebookAppID string
}


type AppContext struct {
	Meta        		Meta
	Localizer   		*i18n.Localizer
	CurrentPath 		string
	CurrentLocale 	string
	Lang 						string
	SEO 						SEOData
	MarkdownContent map[string]map[string]templ.Component
	JSONLDScript    string
	// Add any other shared data here
}