package main

import (
	"os"
	"os/exec"
	"runtime"
)

func main() {
	url := "http://localhost:5000"

	// Check the operating system
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	case "windows":
		exec.Command("cmd", "/c", "start", url).Start()
	default:
		println("Unsupported operating system")
		os.Exit(1)
	}
}
