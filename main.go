package main

import (
	"YoanyWoany/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	a := app.New()
	w := a.NewWindow("Yoany Woany")
	w.Resize(fyne.NewSize(800, 600))

	w.SetContent(ui.BuildStartMenu())

	w.ShowAndRun()

}
