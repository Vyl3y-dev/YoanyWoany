package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildAnotherWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("YOANY WOANY")
	w.Resize(fyne.NewSize(800, 600))
	w.SetContent(AnotherWindowItems(a))
	return w
}

func AnotherWindowItems(a fyne.App) fyne.CanvasObject {
	anotherLabel := widget.NewLabel("OMG! It is another window!")
	anotherBtn := widget.NewButton("Press to go Back!", func() {
		BuildStartMenu(a).Show()
	})
	itemLayout := container.NewWithoutLayout(anotherLabel, anotherBtn)
	anotherLabel.Move(fyne.NewPos(100, 50))
	anotherBtn.Move(fyne.NewPos(100, 100))
	anotherBtn.Resize(fyne.NewSize(150, 50))
	return itemLayout
}
