package cmd

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func LongRunningHandler(sesseion *discordgo.Session, interaction *discordgo.InteractionCreate) {
	//log user interaction: NAME: <username>#<discriminator> ID: <user id> COMMAND: <command name>
	//log.Printf("NAME: %s ID: %s COMMAND: %s", interaction.User.Username, interaction.User.ID, "long")
	sesseion.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	// do your stuff here
	// sleep represents work being done
	time.Sleep(10 * time.Second)

	sesseion.FollowupMessageCreate(interaction.Interaction, true, &discordgo.WebhookParams{
		Content: "Pretend this is an Image of some b-field thingy",
	})
}
