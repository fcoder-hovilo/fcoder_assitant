// scheduler/scheduler.go
// change this scheduler to google calendar api

package scheduler

import (
	"fmt"
	"time"

	"fcoder_assitant/config"

	"github.com/bwmarrin/discordgo"
)

// Start initializes and starts the periodic message scheduler
func Start(s *discordgo.Session, config config.Config) {
	// Schedule the periodic message every 1 minute
	go schedulePeriodicMessage(s, config)
}

func schedulePeriodicMessage(s *discordgo.Session, config config.Config) {
	// Set the time for the first periodic message (1 minute from now)
	firstTime := time.Now().Add(time.Minute)
	periodicTimer := time.NewTimer(time.Until(firstTime))

	for {
		select {
		case <-periodicTimer.C:
			// Send a message to all channels
			sendPeriodicMessage(s, config)

			// Reset the timer for the next 1 minute
			periodicTimer.Reset(time.Minute)
		}
	}
}

func sendPeriodicMessage(s *discordgo.Session, config config.Config) {
	// Replace this message with the desired periodic message
	message := "This is a periodic message sent every 1 minute. asdfasdfasd"

	// Get a list of all guilds the bot is a member of
	guilds, err := s.UserGuilds(100, "", "")
	if err != nil {
		fmt.Println("Error getting guilds:", err)
		return
	}

	// Iterate through each guild
	for _, guild := range guilds {
		// Iterate through each channel in the guild
		channels, err := s.GuildChannels(guild.ID)
		if err != nil {
			fmt.Println("Error getting channels for guild", guild.Name, ":", err)
			continue
		}

		for _, channel := range channels {
			// Send the message to each text channel
			if channel.Type == discordgo.ChannelTypeGuildText {
				s.ChannelMessageSend(channel.ID, message)
			}
		}
	}
}
