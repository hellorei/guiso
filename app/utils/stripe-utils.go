package utils

import "os"

func GetStripeLink(linkName string) string {
	
	link := os.Getenv(linkName)

	if link == "" {
		link = os.Getenv("STRIPE_URL")
	}

	if link == "" {
		return "/"
	}
	
	return link
}