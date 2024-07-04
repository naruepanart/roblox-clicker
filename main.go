package main

import (
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
)

var (
	autoClicking bool
	logLabel     *widget.Label
)

func main() {
	a := app.New()
	w := a.NewWindow("Roblox Auto Click by nptfreez 1.0.0")
	// Resize the window
	w.Resize(fyne.NewSize(480, 100))
	// Create log label
	logLabel = widget.NewLabel("Status: Ready")

	// Buttons and their functionality
	startButton := widget.NewButton("Start (F1)", func() {
		startAutoClick()
	})
	stopButton := widget.NewButton("Stop (F2)", func() {
		stopAutoClick()
	})
	exitButton := widget.NewButton("Exit (F3)", func() {
		stopAutoClick()
		a.Quit()
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
			a.Quit()
		}
	})

	// Layout using VBox container
	w.SetContent(container.NewVBox(
		logLabel,
		startButton,
		stopButton,
		exitButton,
	))

	w.ShowAndRun()
}

func startAutoClick() {
	if autoClicking {
		updateLog("Status: running already")
		return
	}
	autoClicking = true
	go autoClick()
	updateLog("Status: started")
}

func stopAutoClick() {
	if !autoClicking {
		return
	}
	autoClicking = false
	updateLog("Status: stopped")
	// Delay for 5 seconds (adjust as needed)
	time.Sleep(5 * time.Second)
	// Status: ready
	updateLog("Status: ready")
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

func updateLog(message string) {
	logLabel.SetText(message)
}
