package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (A *App) startup(ctx context.Context) {
	A.ctx = ctx
}

// Greet returns a greeting for the given name
func (A *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (A *App) Test() {

}

func (A *App) shutdown(ctx context.Context) {
	fmt.Println("shutting down...")
}
