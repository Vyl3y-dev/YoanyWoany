package main

import (
	"YoanyWoany/windows"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	a := app.New()
	mainWindow := a.NewWindow("YOANYWOANY")
	mainWindow.Resize(fyne.NewSize(1280, 768))

	mainWindow.SetContent(windows.DisplayContent(mainWindow))

	mainWindow.ShowAndRun()

}
