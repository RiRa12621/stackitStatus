package main

import (
	"os"
	"stackitStatus/pkg"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	err := pkg.RunTray()
	if err != nil {
		os.Exit(1)
	}
	a.Run()
}
