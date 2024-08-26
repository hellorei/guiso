package utils

import (
	"encoding/json"
	"fmt"
	"guiso/app/types"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/piprate/json-gold/ld"
)

func GetDefaultSEO(localizer *i18n.Localizer, baseURL string) types.SEOData {
	return types.SEOData{
		Title:       localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Meta.Title"}),
		Description: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Meta.Description"}),
		URL:         baseURL,
		Image:       baseURL + "/images/og-image.jpg",
	}
}

func UpdateSEO(seo types.SEOData, updates map[string]string) types.SEOData {
	if title, ok := updates["title"]; ok {
		seo.Title = title
	}
	if description, ok := updates["description"]; ok {
		seo.Description = description
	}
	if url, ok := updates["url"]; ok {
		seo.URL = url
	}
	if image, ok := updates["image"]; ok {
		seo.Image = image
	}
	return seo
}

func GetJSONLD(appCtx types.AppContext) string {
	return `<script type="application/ld+json">{
		"@context": "https://schema.org",
		"@type": "Organization",
		"name": { ` + appCtx.SEO.Title + ` },
		"url": { ` + appCtx.SEO.URL + ` },
		"logo": { ` + appCtx.SEO.Image + ` },
		"sameAs": [
			` + strings.Join(GetSocialLinksURLs(), ",") + `
		]
	}</script>`
}

type JSONLDOrganization struct {
	Context    string   `json:"@context"`
	Type       string   `json:"@type"`
	Name       string   `json:"name"`
	URL        string   `json:"url"`
	Logo       string   `json:"logo"`
	SameAs     []string `json:"sameAs"`
}

// func GenerateJSONLD(seo types.SEOData) (string) {
// 	org := JSONLDOrganization{
// 		Context: "https://schema.org",
// 		Type:    "Organization",
// 		Name:    seo.Title,
// 		URL:     seo.URL,
// 		Logo:    seo.Image,
// 		SameAs:  GetSocialLinksURLs(),
// 	}

// 	jsonData, err := json.Marshal(org)
// 	if err != nil {
// 		return ""
// 	}

// 	return fmt.Sprintf("<script type=\"application/ld+json\">%s</script>", string(jsonData))
// }

func GenerateJSONLD(localizer *i18n.Localizer) (string, error) {
	doc := map[string]interface{}{
		"@context": "https://schema.org",
		"@type": []string{"Organization", "NonProfit"},
		"name": localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Meta.Organization"}),
		"url": "https://guiso.io",
		"logo": "https://guiso.io/images/og-image.jpg",
		"sameAs": GetSocialLinksURLs(),
		"description": localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Meta.Description"}),
		"event": map[string]interface{}{
				"@type": "FundraiserEvent",
				"name": localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Meta.Title"}),
				"description": localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Meta.Description"}),
		},
	}

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	expanded, err := proc.Expand(doc, options)
	if err != nil {
		return "", err
	}

	compacted, err := proc.Compact(expanded, nil, options)
	if err != nil {
		return "", err
	}

	jsonLD, err := json.MarshalIndent(compacted, "", "  ")
	if err != nil {
		return "", err
	}

	// Escape the JSON and wrap it in a script tag
	scriptTag := fmt.Sprintf("<script type=\"application/ld+json\">\n%s\n</script>", string(jsonLD))

	return scriptTag, nil
}