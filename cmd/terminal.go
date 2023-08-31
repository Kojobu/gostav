package cmd

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func TerminalHander(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	args := []string{}

	if len(options) > 2 {
		if option, ok := optionMap["passw"]; ok {
			args = append(args, option.StringValue())
		}

		if option, ok := optionMap["prog"]; ok {
			args = append(args, option.StringValue())
		}

		if option, ok := optionMap["addargs"]; ok {
			args = append(args, option.StringValue())
		}

		cmds := strings.Split(args[2], " ")

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: terminal(args[1], cmds, args[0]),
			},
		})
	} else {
		if option, ok := optionMap["passw"]; ok {
			args = append(args, option.StringValue())
		}

		if option, ok := optionMap["prog"]; ok {
			args = append(args, option.StringValue())
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: terminal2(args[1], args[0]),
			},
		})
	}
}
