package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func StartButton() fyne.CanvasObject {
	startBtn := widget.NewButton("Press to Start!", func() {

	})

	return startBtn
}

func AnotherButton() fyne.CanvasObject {
	anotherBtn := widget.NewButton("Press for another window!", func() {

	})
	return anotherBtn
}
