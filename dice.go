package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/internal/widget"
)

func main() {
	//new app
	a := app.New()
	//new window
	w := a.NewWindow("Dice Game")
	// window, _ := walk.NewMainWindow()

	//resize window
	// window.SetWidth(400)
	// window.SetHeight(200)
	// w.Resize(fyne.Size(400, 400))

	//image of a dice
	img := canvas.NewImageFromFile("img/dice6.jpg")
	img.FillMode = canvas.ImageFillOriginal

	//button
	btn := widget.NewButton("Play", func() {
		//logic
		rand := rand.Intn(6) + 1
		img.File = fmt.Sprintf("img/dice%d.jpg", rand)
		img.Refresh()
	})
	//setup content and finish Ui
	w.SetContent(
		//NewVBox.. Morethan on widgets
		container.NewVBox(
			img,
			btn,
		),
	)

	//show and run app
	w.ShowAndRun()
}
