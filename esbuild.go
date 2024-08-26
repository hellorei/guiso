package main

import (
	"log"

	"github.com/evanw/esbuild/pkg/api"
)

func buildJS() {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"app/ts/main.ts"},
		Bundle:      true,
		Write:       true,
		Outfile:     "public/js/dist/bundle.js",
		Target:      api.ES2015,
	})

	if len(result.Errors) > 0 {
		log.Fatal("Error building JS")
	}
}
