package main

import (
	"api-client/src/configuration"
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

	app := NewApp()
	xdgUserDir := configuration.NewXDG()
	configurationReadWriter := configuration.NewReadWriter(xdgUserDir)

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "Api-Client",
		Frameless: true,
		Width:     500,
		Height:    500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
			configurationReadWriter,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
