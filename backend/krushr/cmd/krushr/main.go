package main

import (
	"github.com/stanhoenson/krushr/internal/app"
)

func main() {
	var newApp, err = app.CreateApp()
	if err != nil {
		panic("couldn't create app")
	}
	app.Initialize(newApp)
}
