package cmd

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func ImpressumHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Printf("NAME: %s ID: %s COMMAND: %s", i.User.Username, i.User.ID, "impressum")
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: impressum(),
		},
	})
}
