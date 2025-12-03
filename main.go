package main

import (
	"YoanyWoany/db"
	"YoanyWoany/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	a := app.New()
	mainWindow := a.NewWindow("YOANYWOANY")

	db.Connect()
	accountsPage := ui.AccountsPage(mainWindow)

	mainWindow.SetContent(accountsPage)

	mainWindow.Resize(fyne.NewSize(450, 450))

	mainWindow.ShowAndRun()

}
