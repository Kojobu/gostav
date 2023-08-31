package cmd

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func LongRunningHandler(sesseion *discordgo.Session, interaction *discordgo.InteractionCreate) {
	sesseion.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	// do your stuff here
	// sleep represents work being done
	time.Sleep(10 * time.Second)

	sesseion.FollowupMessageCreate(interaction.Interaction, true, &discordgo.WebhookParams{
		Content: "Done! this took 10 seconds",
	})
}
