package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func BFieldHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Printf("NAME: %s ID: %s COMMAND: %s", i.User.Username, i.User.ID, "b-field")
	dat_path := "/home/potato/Documents/sensorsave.dat"
	img_path := b_plot(dat_path)
	f, err := os.Open(img_path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Here's the local B-field of my room:",
			Files: []*discordgo.File{
				&discordgo.File{
					Name:   img_path,
					Reader: f,
				},
			},
		},
	})
}
