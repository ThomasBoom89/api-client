package main

import (
	"api-client/src/app"
	"api-client/src/configuration"
	"api-client/src/database"
	"api-client/src/frontend"
	"api-client/src/runtime"
	"embed"
	"github.com/rs/zerolog"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if true {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	context := app.NewContext()
	xdgUserDir := configuration.NewXDG()
	databaseClient := database.NewClient(xdgUserDir)
	database.AutoMigrate(databaseClient)
	configurationReadWriter := configuration.NewReadWriter(xdgUserDir)
	config := frontend.NewConfiguration(configurationReadWriter)
	projectRepository := database.NewRepository[database.Project](databaseClient)
	collectionsRepository := database.NewRepository[database.Collection](databaseClient)
	httpRequestRepository := database.NewHttpRequestRepository(databaseClient)
	websocketRequestRepository := database.NewWebsocketRequestRepository(databaseClient)
	request := frontend.NewRequest(httpRequestRepository, configurationReadWriter)
	wailsEvent := runtime.Wails{}
	websocket := frontend.NewWebsocket(context, &wailsEvent)
	projects := frontend.NewProjects(projectRepository)
	collections := frontend.NewCollections(collectionsRepository)
	httpRequest := frontend.NewHttpRequests(httpRequestRepository)
	websocketRequest := frontend.NewWebsocketRequests(websocketRequestRepository)
	requests := frontend.NewRequests(httpRequest, websocketRequest)
	// Create application with options
	err := wails.Run(&options.App{
		Title:     "Api-Client",
		Width:     960,
		Height:    720,
		MaxWidth:  7680, // 8k
		MaxHeight: 4320,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        context.SetContext,
		OnShutdown:       context.Cancel,
		Bind: []interface{}{
			config,
			request,
			projects,
			collections,
			requests,
			httpRequest,
			websocketRequest,
			websocket,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
