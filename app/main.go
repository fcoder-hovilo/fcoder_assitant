// main.go

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"fcoder_assitant/config"
	"fcoder_assitant/handlers"
	"fcoder_assitant/scheduler"

	"github.com/bwmarrin/discordgo"
)

func main() {
	config, err := config.Load()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Initialize and start scheduler
	scheduler.Start(dg, config)

	// Initialize and register message handler
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		handlers.MessageCreate(s, m, config)
	})

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	// Wait for a CTRL+C signal to gracefully close the bot
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Close the Discord session cleanly before exiting
	dg.Close()
}
