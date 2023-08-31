package cmd

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func BasicFileHander(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//log.Printf("NAME: %s ID: %s COMMAND: %s", i.User.Username, i.User.ID, "basic-file")
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey there! Congratulations, you just executed your first slash command with a file in the response",
			Files: []*discordgo.File{
				{
					ContentType: "text/plain",
					Name:        "test.txt",
					Reader:      strings.NewReader("Hello Discord!!"),
				},
			},
		},
	})
}
