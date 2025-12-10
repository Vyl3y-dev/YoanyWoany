package windows

import (
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
	)
	return displayContent
}

func OpenPortal(w fyne.Window) *fyne.Container {

	newLabel := widget.NewLabel("Press to open the portal")
	newLabel.Move(fyne.NewPos(700, 200))
	newLabel.Resize(fyne.NewSize(200, 100))
	portalButton := widget.NewButton("Open Portal", func() { dialog.ShowInformation("Portal Status", "A Portal Has Opened", w) })
	portalButton.Move(fyne.NewPos(700, 300))
	portalButton.Resize(fyne.NewSize(200, 100))

	displayContent := container.NewWithoutLayout(
		newLabel,
		portalButton,
	)

	return displayContent
}

func HiddenFacts() *fyne.Container {
	forbiddenFactLabel := widget.NewLabel("You wanna know a forbidden Fact?")
	forbiddenFactLabel.Move(fyne.NewPos(700, 400))
	forbiddenFactLabel.Resize(fyne.NewSize(200, 100))

	forbiddenFacts := widget.NewButton("Forbidden Fact", func() { forbiddenFactLabel.SetText("Cats know the secrets of man!") })
	forbiddenFacts.Move(fyne.NewPos(700, 500))
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

	size := float32(512) // 20–60 px
	x := float32(100)
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
	t.Move(fyne.NewPos(1000, 100))
	t.Resize(fyne.NewSize(150, 150))

	if isCursed == true {
		t.Text = "There exists a cursed notebook"
		t.Refresh()
	} else {

		t.Text = "There does not exist a cursed notebook"
		t.Refresh()
	}

	boxForText := container.NewWithoutLayout(t)

	return boxForText
}
