package cmd

import (
	"github.com/bwmarrin/discordgo"
)

func TestHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//log.Printf("NAME: %s ID: %s COMMAND: %s", i.User.Username, i.User.ID, "test")
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Test!",
		},
	})
}
