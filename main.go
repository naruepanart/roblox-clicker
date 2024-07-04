package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
)

var (
	stopChan chan struct{}
)

func main() {
	a := app.New()
	w := a.NewWindow("Roblox Auto Click by nptfreez")

	// Create buttons and set their functionality
	startButton := widget.NewButton("Start (F1)", func() {
		startAutoClick()
	})
	stopButton := widget.NewButton("Stop (F2)", func() {
		stopAutoClick()
	})

	// Define keyboard shortcuts
	w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeyF1:
			startAutoClick()
		case fyne.KeyF2:
			stopAutoClick()
		case fyne.KeyF3:
			stopAutoClick()
			a.Quit() // Quit application on F3 press
		}
	})

	// Layout using VBox container
	w.SetContent(container.NewVBox(
		widget.NewLabel("Roblox Auto Click by nptfreez"),
		startButton,
		stopButton,
	))

	w.ShowAndRun()
}

func startAutoClick() {
	if stopChan != nil {
		fmt.Println("Auto Click already running")
		return
	}
	stopChan = make(chan struct{})
	go autoClick()
	fmt.Println("Auto Click started.")
}

func stopAutoClick() {
	if stopChan != nil {
		close(stopChan)
		stopChan = nil
		fmt.Println("Auto Click stopped")
	}
}

func autoClick() {
	for {
		select {
		case <-stopChan:
			fmt.Println("Auto Click stopped")
			return
		default:
			// Perform left click
			robotgo.Click("left")

			// Delay for 500 milliseconds (adjust as needed)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
