package utils

import (
	"encoding/xml"
	"time"
)

type URL struct {
	Loc        string    `xml:"loc"`
	LastMod    time.Time `xml:"lastmod"`
	ChangeFreq string    `xml:"changefreq"`
	Priority   float32   `xml:"priority"`
}

type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

func GenerateSitemap(baseURL string) ([]byte, error) {
	urls := []URL{
		{Loc: baseURL + "/", LastMod: time.Now(), ChangeFreq: "daily", Priority: 1.0},
		{Loc: baseURL + "/about", LastMod: time.Now(), ChangeFreq: "weekly", Priority: 0.8},
		{Loc: baseURL + "/situation/persecution", LastMod: time.Now(), ChangeFreq: "weekly", Priority: 0.7},
		{Loc: baseURL + "/situation/censorship", LastMod: time.Now(), ChangeFreq: "weekly", Priority: 0.7},
		{Loc: baseURL + "/situation/violence", LastMod: time.Now(), ChangeFreq: "weekly", Priority: 0.7},
		{Loc: baseURL + "/get-involved", LastMod: time.Now(), ChangeFreq: "weekly", Priority: 0.9},
		{Loc: baseURL + "/contact", LastMod: time.Now(), ChangeFreq: "monthly", Priority: 0.6},
	}

	sitemap := Urlset{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}

	return xml.MarshalIndent(sitemap, "", "  ")
}