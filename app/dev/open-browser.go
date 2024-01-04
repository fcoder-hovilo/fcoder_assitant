package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	url := "http://localhost:5000"

	// Check the operating system
	switch runtime.GOOS {
	case "linux":
		dockerCmd := exec.Command("sudo", "docker-compose", "up", "--build", "-d")
		if err := runAndWait(dockerCmd); err != nil {
			fmt.Println("Error running Docker command:", err)
			os.Exit(1)
		}
	case "darwin":
		// Assume macOS uses the same command as Linux
		dockerCmd := exec.Command("docker-compose", "up", "--build", "-d")
		if err := runAndWait(dockerCmd); err != nil {
			fmt.Println("Error running Docker command:", err)
			os.Exit(1)
		}
	case "windows":
		dockerCmd := exec.Command("docker-compose", "up", "--build", "-d")
		if err := runAndWait(dockerCmd); err != nil {
			fmt.Println("Error running Docker command:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unsupported operating system")
		os.Exit(1)
	}

	// Wait for 10 seconds
	time.Sleep(10 * time.Second)

	// Open the browser
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	case "windows":
		exec.Command("cmd", "/c", "start", url).Start()
	default:
		fmt.Println("Unsupported operating system")
		os.Exit(1)
	}
}

func runAndWait(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
