package main

import (
	"YoanyWoany/ui"

	"fyne.io/fyne/v2/app"
)

func main() {

	a := app.New()
	mainWindow := ui.BuildStartMenu(a)

	mainWindow.ShowAndRun()

}
