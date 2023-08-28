package main

import (
	app "RestJwtAuth/internal/pkg/app"
	"log"
)

const (
	ADDRESS = ":8080"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
}
