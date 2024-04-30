package main

import (
	"context"
	"debugger-api/internal/app"
	"log"

	_ "debugger-api/docs"
	_ "github.com/swaggo/http-swagger/v2"
)

// @title Debugger api
// @version 1.0
// @description Api for http debugging
// @BasePath /api/v1

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	log.Println("Running server")
	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
