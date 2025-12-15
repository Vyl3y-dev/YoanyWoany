package windows

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func DisplayContent(w fyne.Window) *fyne.Container {

	displayContent := container.NewWithoutLayout(
		OpenPortal(w),
		HiddenFacts(),
		FloatingMug(),
		CursedNotebook(),
		RunningChair(),
	)
	return displayContent
}

func OpenPortal(w fyne.Window) *fyne.Container {

	newLabel := widget.NewLabel("Press to open the portal")
	newLabel.Move(fyne.NewPos(500, 200))
	newLabel.Resize(fyne.NewSize(200, 100))
	portalButton := widget.NewButton("Open Portal", func() { dialog.ShowInformation("Portal Status", "A Portal Has Opened", w) })
	portalButton.Move(fyne.NewPos(500, 300))
	portalButton.Resize(fyne.NewSize(200, 100))

	displayContent := container.NewWithoutLayout(
		newLabel,
		portalButton,
	)

	return displayContent
}

func HiddenFacts() *fyne.Container {
	forbiddenFactLabel := widget.NewLabel("You wanna know a forbidden Fact?")
	forbiddenFactLabel.Move(fyne.NewPos(500, 400))
	forbiddenFactLabel.Resize(fyne.NewSize(200, 100))

	forbiddenFacts := widget.NewButton("Forbidden Fact", func() { forbiddenFactLabel.SetText("Cats know the secrets of man!") })
	forbiddenFacts.Move(fyne.NewPos(500, 500))
	forbiddenFacts.Resize(fyne.NewSize(200, 100))

	displayContent := container.NewWithoutLayout(
		forbiddenFactLabel,
		forbiddenFacts,
	)

	return displayContent
}

func FloatingMug() fyne.CanvasObject {
	img := canvas.NewImageFromFile("assets/YoanyMug.png")
	img.FillMode = canvas.ImageFillOriginal

	size := float32(384) // 20–60 px
	x := float32(0)
	y := float32(100)
	speedY := float32(2)

	img.Resize(fyne.NewSize(size, size))
	img.Move(fyne.NewPos(x, y))

	ticker := time.NewTicker(time.Millisecond * 16)
	go func() {
		for range ticker.C {
			y += speedY
			if y >= 200 || y <= 0 {
				speedY = -speedY
			}

			fyne.Do(func() {
				img.Move(fyne.NewPos(x, y))
				img.FillMode = canvas.ImageFillOriginal
				img.Refresh()
			})
		}
	}()

	return img
}

func CursedNotebook() *fyne.Container {
	isCursed := true
	t := canvas.NewText("", color.White)
	t.TextSize = 24
	t.Move(fyne.NewPos(800, 0))
	t.Resize(fyne.NewSize(150, 150))

	if isCursed == true {
		t.Text = "There exists a cursed notebook"

		t.Refresh()
	} else {

		t.Text = "There does not exist a cursed notebook"
		t.Refresh()
	}

	img := canvas.NewImageFromFile("assets/CursedNotebook.png")
	img.FillMode = canvas.ImageFillOriginal
	size := float32(384) // 20–60 px
	x := float32(800)
	y := float32(100)

	img.Resize(fyne.NewSize(size, size))
	img.Move(fyne.NewPos(x, y))

	newBox := container.NewWithoutLayout(t, img)

	return newBox
}

func SuspiciousPeach() {

}

func RunningChair() *fyne.Container {
	now := time.Now()

	dinnerTime := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		18, 0, 0, 0,
		now.Location(),
	)

	runningChair := widget.NewLabel("The chair at the dinner table plans to run away!")
	runningChair.Move(fyne.NewPos(800, 700))
	runningChair.Resize(fyne.NewSize(150, 150))

	if now.Before(dinnerTime) {
		fmt.Println("Not dinner yet!")
		runningChair.SetText("Not dinner time yet! You will have to wait.")
	} else if now.After(dinnerTime) {
		fmt.Println("Dinner is over 😭")
		runningChair.SetText("You missed dinner time (again)! You will have to wait until tomorrow.")
	} else if now.Equal(dinnerTime) {
		fmt.Println("It's dinner time RIGHT NOW 🍽️")
		runningChair.SetText("Its dinner time RIGHT NOW! Make your escape, while your owner is distracted. HURRY! You only get one chance.")
	}

	newBox := container.NewWithoutLayout(runningChair)

	return newBox
}
