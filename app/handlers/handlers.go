// handlers/handlers.go

package handlers

import (
	"strings"

	"fcoder_assitant/config"

	"github.com/bwmarrin/discordgo"
)

// MessageCreate is the handler for the MessageCreate event
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate, config config.Config) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, config.Prefix) {
		args := strings.Fields(m.Content[len(config.Prefix):])

		switch args[0] {
		case "hello":
			s.ChannelMessageSend(m.ChannelID, "Hello, "+m.Author.Mention()+"!")
		}
	}
}
