package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildStartMenu(a fyne.App) fyne.Window {
	w := a.NewWindow("YOANY WOANY")
	w.Resize(fyne.NewSize(800, 600))
	w.SetContent(StartMenuItems(a))
	return w
}

func StartMenuItems(a fyne.App) fyne.CanvasObject {
	ily := widget.NewLabel("Yoany I love you! ❤️")
	startBtn := widget.NewButton("Press to Start!", func() {
		BuildAnotherWindow(a).Show()
	})
	itemLayout := container.NewWithoutLayout(ily, startBtn)
	ily.Move(fyne.NewPos(100, 50))
	startBtn.Move(fyne.NewPos(100, 100))
	startBtn.Resize(fyne.NewSize(150, 50))
	return itemLayout
}
