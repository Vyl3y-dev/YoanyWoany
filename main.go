package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	mainWindow := a.NewWindow("YOANYWOANY")
	mainWindow.Resize(fyne.NewSize(1280, 768))

	currentContent := OpenPortal(mainWindow)

	mainWindow.SetContent(currentContent)

	mainWindow.ShowAndRun()

}

func OpenPortal(w fyne.Window) *fyne.Container {

	newLabel := widget.NewLabel("Press to open the portal")
	newLabel.Move(fyne.NewPos(500, 200))
	newLabel.Resize(fyne.NewSize(200, 100))
	portalButton := widget.NewButton("Open Portal", func() { dialog.ShowInformation("Portal Status", "A Portal Has Opened", w) })
	portalButton.Move(fyne.NewPos(500, 300))
	portalButton.Resize(fyne.NewSize(200, 100))

	forbiddenFactLabel := widget.NewLabel("You wanna know a forbidden Fact?")
	forbiddenFactLabel.Move(fyne.NewPos(500, 400))
	forbiddenFactLabel.Resize(fyne.NewSize(200, 100))

	forbiddenFacts := widget.NewButton("Forbidden Fact", func() { forbiddenFactLabel.SetText("Cats know the secrets of man!") })
	forbiddenFacts.Move(fyne.NewPos(500, 500))
	forbiddenFacts.Resize(fyne.NewSize(200, 100))

	displayContent := container.NewWithoutLayout(
		newLabel,
		portalButton,
		forbiddenFactLabel,
		forbiddenFacts,
	)

	return displayContent
}
