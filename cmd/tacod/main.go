package main

import (
	"log"

	"github.com/justinas/alice"
	"github.com/sul-dlss-labs/taco"
	"github.com/sul-dlss-labs/taco/config"
	"github.com/sul-dlss-labs/taco/generated/restapi"
	"github.com/sul-dlss-labs/taco/handlers"
	"github.com/sul-dlss-labs/taco/logger"
)

func main() {
	rt, err := taco.NewRuntime(config.NewConfig())
	if err != nil {
		log.Fatalln(err)
	}

	server := createServer(rt)
	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func createServer(rt *taco.Runtime) *restapi.Server {
	api := handlers.BuildAPI(rt)
	server := restapi.NewServer(api)
	handler := alice.New(
		logger.NewRequestLoggerMW(),
	).Then(api.Serve(nil))
	server.SetHandler(handler)
	defer server.Shutdown()

	// set the port this service will be run on
	server.Port = rt.Config().Port
	return server
}
