// handlers/handlers.go

package handlers

import (
	"strings"

	"fcoder_assitant/config"

	"github.com/bwmarrin/discordgo"
)

// MessageCreate is the handler for the MessageCreate event
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate, config config.Config) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if the message starts with the specified prefix
	if strings.HasPrefix(m.Content, config.Prefix) {
		// Split the message into command and arguments
		args := strings.Fields(m.Content[len(config.Prefix):])

		// Check the command and perform the corresponding action
		switch args[0] {
		case "hello":
			// Respond to the "!hello" command
			s.ChannelMessageSend(m.ChannelID, "Hello, "+m.Author.Mention()+"!")
		}
	}
}
