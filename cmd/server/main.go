package main

import "fmt"

//App - contains all pointers to the app like database connection
type App struct{}

//Run - sets up our app
func (app *App) Run() error {
	fmt.Println("Setting up our app")
	return nil
}

func main() {
	fmt.Println("Go rest API")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error start App")
		fmt.Println(err)
	}
}
