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
// @description Api for http debugging
// @BasePath /api/v1

func main() {

	//	var dev map[string]interface{}
	//	err := json.Unmarshal([]byte(`{
	//  "id": 123,
	//  "name": "GoLinux Cloud",
	//  "address": {
	//    "street": "Summer",
	//    "city": "San Jose"
	//  },
	//  "phoneNumber": 1234567890,
	//  "role": "Admin",
	//  "someField": "unstructed"
	//}`), &dev)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Println(dev)
	//
	//	str, _ := json.Marshal(dev)
	//
	//	fmt.Println(string(str))

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
