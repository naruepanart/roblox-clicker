package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
)

const (
	StatusReady   = "Ready"
	StatusStarted = "Started"
	StatusRunning = "Running.."
	StatusStopped = "Stopped"
)

var (
	isAutoClicking bool
	statusLabel    = widget.NewLabel("Status: " + StatusReady)
)

func main() {
	a := app.New()
	w := a.NewWindow("Roblox Auto Clicker 1.0.0")
	icon, _ := fyne.LoadResourceFromPath("icons.svg")
	w.SetIcon(icon)
	w.Resize(fyne.NewSize(480, 100))

	startButton := widget.NewButton("Start (F1)", startAutoClick)
	stopButton := widget.NewButton("Stop (F2)", stopAutoClick)
	exitButton := widget.NewButton("Exit (F3)", a.Quit)

	w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		handleShortcuts(event, a)
	})

	w.SetContent(container.NewVBox(statusLabel, startButton, stopButton, exitButton))
	w.ShowAndRun()
}

func handleShortcuts(event *fyne.KeyEvent, a fyne.App) {
	switch event.Name {
	case fyne.KeyF1:
		startAutoClick()
	case fyne.KeyF2:
		stopAutoClick()
	case fyne.KeyF3:
		a.Quit()
	}
}

func startAutoClick() {
	if isAutoClicking {
		updateStatus("Status: " + StatusRunning)
		return
	}
	isAutoClicking = true
	updateStatus("Status: " + StatusStarted)
	go autoClick()
}

func stopAutoClick() {
	if isAutoClicking {
		isAutoClicking = false
		updateStatus("Status: " + StatusStopped)
		time.Sleep(5 * time.Second)
		updateStatus("Status: " + StatusReady)
	}
}

func autoClick() {
	for isAutoClicking {
		if robotgo.GetTitle() == "Roblox" {
			robotgo.Click("left")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func updateStatus(message string) {
	statusLabel.SetText(message)
}
