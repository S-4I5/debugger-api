package main

import (
	"context"
	_ "debugger-api/docs"
	"debugger-api/internal/app"
	"debugger-api/internal/config"
	_ "github.com/swaggo/http-swagger/v2"
	"log"
)

// @title Debugger api
// @version 1.0
// @description Mock api for http debugging
// @BasePath /api/v1
func main() {
	ctx := context.Background()

	cfg := config.MustLoad("./config/config.yaml")

	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	log.Println("Running server")
	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
