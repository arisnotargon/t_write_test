package main

import (
	"github.com/arisnotargon/t_write_test/app"
)

func main() {
	factoryApp := app.NewApp()

	factoryApp.Run()
}
