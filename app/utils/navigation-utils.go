package utils

import (
	"guiso/app/types"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GetLocalizedMainMenu(localizer *i18n.Localizer) []types.MenuItem {
	return []types.MenuItem{
		{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.OurWork"}), Href: "about"},
		{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.Situation"}), Href: "#", SubMenu: []types.MenuItem{
			{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.SituationPersecution"}), Href: "/situation/persecution", Icon: "/images/elements/persecution.svg#persecution-icon"},
			{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.SituationCensorship"}), Href: "/situation/censorship", Icon: "/images/elements/censorship.svg#censorship-icon"},
			{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.SituationViolence"}), Href: "/situation/violence", Icon: "/images/elements/violence.svg#violence-icon"},
		}},
		{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.GetInvolved"}), Href: "/get-involved"},
	}
}

func GetLocalizedFooterMenu(localizer *i18n.Localizer) []types.MenuItem {
	return []types.MenuItem{
		{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.Home"}), Href: "/"},
		{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.Persecution"}), Href: "/situation/persecution"},
		{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.Censorship"}), Href: "/situation/censorship"},
		{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.Violence"}), Href: "/situation/violence"},
		{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.GetInvolved"}), Href: "/get-involved"},
		{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.About"}), Href: "/about"},
		{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.Contact"}), Href: "/contact"},
		{Label: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MenuLabels.Privacy"}), Href: "/privacy-policy"},
	}
}

func GetSocialLinks() []types.SocialLink {
	return []types.SocialLink{
		// {Label: "Facebook", Href: "#", Icon: "/images/icons/facebook.svg#facebook-icon"},
		{
			Label: "X",
			Href: "https://twitter.com/guiso",
			Handler: "@guiso",
			Icon: "/images/icons/twitterx.svg#twitterx-icon",
		},
		{
			Label: "TikTok",
			Handler: "@guiso",
			Href: "https://www.tiktok.com/@guiso",
			Icon: "/images/icons/tiktok.svg#tiktok-icon",
		},
		{
			Label: "Instagram",
			Handler: "@guiso",
			Href: "https://www.instagram.com/guiso/",
			Icon: "/images/icons/instagram.svg#instagram-icon",
		},
		// {Label: "Youtube", Href: "#", Icon: "/images/icons/youtube.svg#youtube-icon"},
	}
}

func GetSocialLinksURLs() []string {
	urls := []string{}
	socialLinks := GetSocialLinks()
	
	for _, socialLink := range socialLinks {
		urls = append(urls, socialLink.Href)
	}

	return urls
}

func GetLanguageMenuItems() []types.MenuItem {
	return []types.MenuItem{
		{Label: "English", Href: "en"},
		{Label: "Español", Href: "es"},
		// {Label: "Français", Href: "fr"},
		// {Label: "Português", Href: "pt"},
	}
}
