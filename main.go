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
	w := a.NewWindow("Auto Clicker")

	// Create buttons and set their functionality
	startButton := widget.NewButton("Start (F2)", func() {
		startAutoClick()
	})
	stopButton := widget.NewButton("Exit (F3)", func() {
		stopAutoClick()
	})

	// Define keyboard shortcuts
	w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		if event.Name == fyne.KeyF2 {
			startAutoClick()
		}
		if event.Name == fyne.KeyF3 {
			stopAutoClick()
			a.Quit() // Quit application on F3 press
		}
	})

	// Layout using VBox container
	w.SetContent(container.NewVBox(
		widget.NewLabel("Roblox Auto Clicker"),
		startButton,
		stopButton,
	))

	w.ShowAndRun()
}

func startAutoClick() {
	if stopChan != nil {
		fmt.Println("Auto clicker already running.")
		return
	}
	stopChan = make(chan struct{})
	go autoClick()
	fmt.Println("Auto clicker started.")
}

func stopAutoClick() {
	if stopChan != nil {
		close(stopChan)
		stopChan = nil
		fmt.Println("Auto clicker stopped.")
	}
}

func autoClick() {
	for {
		select {
		case <-stopChan:
			fmt.Println("Auto clicker stopped.")
			return
		default:
			// Perform left click
			robotgo.Click("left")

			// Delay for 1 second
			time.Sleep(500 * time.Microsecond)
		}
	}
}
