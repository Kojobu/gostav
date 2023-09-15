package cmd

import (
	"github.com/bwmarrin/discordgo"
)

func MensaHander(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//log.Printf("NAME: %s ID: %s COMMAND: %s", i.User.Username, i.User.ID, "mensa")
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: mensa_scrap(3),
		},
	})
}
