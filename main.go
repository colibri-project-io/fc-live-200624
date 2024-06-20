package main

import (
	"github.com/colibri-project-io/colibri-sdk-go"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"github.com/colibri-project-io/fc-live-200624/src"
)

func init() {
	colibri.InitializeApp()
	// cacheDB.Initialize()
	// sqlDB.Initialize()
}

func main() {
	restserver.AddRoutes(src.NewCityController().Routes())

	restserver.ListenAndServe()
}
