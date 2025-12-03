package ui

import (
	"YoanyWoany/db"
	"YoanyWoany/models"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AccountsPage(win fyne.Window) fyne.CanvasObject {

	// 🧁 Input fields
	usernameEntry := widget.NewEntry()
	emailEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()

	// Labels
	usernameLabel := widget.NewLabel("Username:")
	emailLabel := widget.NewLabel("Email:")
	passwordLabel := widget.NewLabel("Password:")

	// 🧁 account list
	accountList := widget.NewLabel("")

	// 🧁 Buttons
	addBtn := widget.NewButton("Add Account", func() {
		username := usernameEntry.Text
		email := emailEntry.Text
		password := passwordEntry.Text

		if username == "" || email == "" || password == "" {
			accountList.SetText("❌ All fields required, kotence!")
			return
		}

		// Create the model instance
		newAccount := models.Account{
			Username: username,
			Email:    email,
			Password: password, // (hash later)
		}

		// Insert into DB
		result := db.DB.Create(&newAccount)
		if result.Error != nil {
			accountList.SetText(fmt.Sprintf("❌ Error: %v", result.Error))
			return
		}

		// Show success
		accountList.SetText(
			fmt.Sprintf("✨ Added: %s (%s)", newAccount.Username, newAccount.Email),
		)

		// Clear fields
		usernameEntry.SetText("")
		emailEntry.SetText("")
		passwordEntry.SetText("")
	})
	searchBtn := widget.NewButton("Search", func() {
		// logic later
	})
	deleteBtn := widget.NewButton("Delete", func() {
		// logic later
	})

	// Layout
	form := container.NewGridWithColumns(2,
		usernameLabel, usernameEntry,
		emailLabel, emailEntry,
		passwordLabel, passwordEntry,
	)

	actions := container.NewHBox(addBtn, searchBtn, deleteBtn)

	page := container.NewVBox(
		widget.NewLabelWithStyle("👤 Accounts Playground", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		form,
		actions,
		widget.NewSeparator(),
		accountList,
	)

	return page
}
