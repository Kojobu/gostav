package cmd

import "github.com/bwmarrin/discordgo"

func MensaHander(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: mensa_scrap(false),
		},
	})
}
