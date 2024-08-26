package utils

import "os"

func GetGoogleAnalyticsMeasurementID() string {
	measurementID := os.Getenv("GA_MEASUREMENT_ID")
	if measurementID == "" {
		return ""
	}
	return measurementID
}

func GetGoogleAnalyticsAPISecret() string {
	apiSecret := os.Getenv("GA_API_SECRET")
	if apiSecret == "" {
		return ""
	}
	return apiSecret
}

func GetGoogleAnalyticsTagScriptURL() string {
	tagScriptURL := "https://www.googletagmanager.com/gtag/js?id=" + GetGoogleAnalyticsMeasurementID()
	return tagScriptURL
}
