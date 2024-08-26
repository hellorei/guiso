package utils

import (
	"net/url"
	"path"
	"strings"
)

// ReplaceLocaleInURL replaces the current locale in the URL with a new one
func ReplaceLocaleInURL(currentURL, newLocale string) (string, error) {
	u, err := url.Parse(currentURL)
	if err != nil {
		return "", err
	}

	path := u.Path
	segments := strings.Split(path, "/")

	if len(segments) > 1 && isLocale(segments[1]) {
		segments[1] = newLocale
	} else {
		segments = append([]string{"", newLocale}, segments[1:]...)
	}

	u.Path = strings.Join(segments, "/")
	return u.String(), nil
}

// GetCurrentLocale extracts the current locale from the URL
func GetCurrentLocale(currentURL string) string {
	u, err := url.Parse(currentURL)
	if err != nil {
		return "" // or a default locale
	}

	segments := strings.Split(u.Path, "/")
	if len(segments) > 1 && isLocale(segments[1]) {
		return segments[1]
	}

	return "" // or a default locale
}

// isLocale checks if a string is a valid locale
func isLocale(s string) bool {
	// Add your locale validation logic here
	// For example, check against a list of supported locales
	supportedLocales := []string{"en", "es", "fr", "jp"} // Add your supported locales
	for _, locale := range supportedLocales {
		if s == locale {
			return true
		}
	}
	return false
}

func LocalizedURL(lang, path string) string {
	// Remove leading slash if present
	path = strings.TrimPrefix(path, "/")
	
	// If path is empty, just return the language
	if path == "" {
		return "/" + lang
	}
	
	// If path already starts with the language code, return as is
	if strings.HasPrefix(path, lang+"/") {
		return "/" + path
	}
	
	// Otherwise, prepend the language code
	return "/" + lang + "/" + path
}

// SwitchLanguageURL changes the language of the current URL
func SwitchLanguageURL(currentURL, newLang string) string {
	parts := strings.SplitN(strings.TrimPrefix(currentURL, "/"), "/", 2)
	if len(parts) == 1 {
		return "/" + newLang
	}
	return "/" + path.Join(newLang, parts[1])
}