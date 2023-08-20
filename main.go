package main

import "github.com/biskitsx/Grocery/app"

func main() {
	if err := app.SetupAndRunApp(); err != nil {
		panic(err)
	}
}
