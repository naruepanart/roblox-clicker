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
	currentStatus  string
)

func main() {
	currentStatus = StatusReady
	fmt.Println("Status:", currentStatus)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Press '1' to start, '2' to stop, and '3' to quit.")

	go handleInput()

	for range signalChannel {
		stopAutoClicking()
		fmt.Println("Exiting...")
		os.Exit(0)
	}
}

func handleInput() {
	var input string
	for {
		fmt.Scanln(&input)
		switch input {
		case "1":
			startAutoClicking()
		case "2":
			stopAutoClicking()
		case "3":
			stopAutoClicking()
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid input. Press '1' to start, '2' to stop, and '3' to quit.")
		}
	}
}

func startAutoClicking() {
	if isAutoClicking {
		fmt.Println("Status:", StatusRunning)
		return
	}
	isAutoClicking = true
	currentStatus = StatusStarted
	fmt.Println("Status:", currentStatus)

	go func() {
		for isAutoClicking {
			if robotgo.GetTitle() == "Roblox" {
				robotgo.Click("left")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

func stopAutoClicking() {
	if isAutoClicking {
		isAutoClicking = false
		currentStatus = StatusStopped
		fmt.Println("Status:", currentStatus)
	}
}
