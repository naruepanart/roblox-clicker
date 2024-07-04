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
	autoClicking bool
)

func main() {
	a := app.New()
	w := a.NewWindow("Roblox Auto Click by nptfreez")
	// Resize the window
	w.Resize(fyne.NewSize(300, 100))
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
	if autoClicking {
		fmt.Println("Auto Click already running")
		return
	}
	autoClicking = true
	go autoClick()
	fmt.Println("Auto Click started")
}

func stopAutoClick() {
	if !autoClicking {
		return
	}
	autoClicking = false
	fmt.Println("Auto Click stopped")
}

func autoClick() {
	for autoClicking {
		// Check if the active window title matches "Roblox" before clicking
		title := robotgo.GetTitle()
		if title == "Roblox" {
			// Perform left click
			robotgo.Click("left")
		}

		// Delay for 500 milliseconds (adjust as needed)
		time.Sleep(500 * time.Millisecond)
	}
}
