package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Счётчик")

	// состояния
	prefs := myApp.Preferences()
	counter := prefs.Int("counter") // получение

	counterLabel := widget.NewLabel(fmt.Sprintf("Счётчик: %d", counter))

	incrementButton := widget.NewButton("Увеличить счётчик", func() {
		counter++
		counterLabel.SetText(fmt.Sprintf("Счётчик: %d", counter))
		prefs.SetInt("counter", counter) // сохранение
	})

	resetButton := widget.NewButton("Обнулить счётчик", func() {
		counter = 0
		counterLabel.SetText(fmt.Sprintf("Счётчик: %d", counter))
		prefs.SetInt("counter", counter) // сохранение
	})

	myWindow.SetContent(container.NewVBox(
		counterLabel,
		incrementButton,
		resetButton,
	))

	myWindow.ShowAndRun()
}
