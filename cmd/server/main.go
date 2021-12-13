package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/commenting-system/internal/transport/http"
)

//App - contains all pointers to the app like database connection
type App struct{}

//Run - sets up our app
func (app *App) Run() error {
	fmt.Println("Setting up our app")

	handler := transportHTTP.NewHandler()

	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go rest API")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting App")
		fmt.Println(err)
	}
}
