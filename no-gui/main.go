package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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
)

func main() {
	fmt.Println("Status:", StatusReady)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch input {
		case "1":
			start()
		case "2":
			stop()
		case "3":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid input. Press '1' to start, '2' to stop, and '3' to quit.")
		}
	}
}

func start() {
	if isAutoClicking {
		fmt.Println("Status:", StatusRunning)
		return
	}
	isAutoClicking = true
	fmt.Println("Status:", StatusStarted)
	go autoClick()
}

func stop() {
	if isAutoClicking {
		isAutoClicking = false
		fmt.Println("Status:", StatusStopped)
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
