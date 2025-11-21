package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	a := app.New()
	mainWindow := a.NewWindow("YOANYWOANY")
	mainWindow.Resize(fyne.NewSize(450, 450))

	mainWindow.ShowAndRun()

}
