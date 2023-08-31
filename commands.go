package main

import "github.com/bwmarrin/discordgo"

var all_commands []*discordgo.ApplicationCommand = []*discordgo.ApplicationCommand{
	{
		Name:        "mensa",
		Description: "Yields the menu of the Mensa INF 306.",
	},

	{
		Name:        "impressum",
		Description: "Reveals the src code and its coder.",
	},

	{
		Name:        "terminal",
		Description: "Execute the given line and yields result.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "passw",
				Description: "Password for authentification.",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "prog",
				Description: "Program executed in the console",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
				// Commands might have choices, think of them like of enum values
			},
			{
				Name:        "addargs",
				Description: "Additional arguments executed after the program",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    false,
				// Commands might have choices, think of them like of enum values
			},
		},
	},
	{
		Name:        "test-command",
		Description: "Testing message returns via functions",
	},

	{
		Name:        "basic-command-with-files",
		Description: "Basic command with files",
	},

	{
		Name:        "b-field",
		Description: "Plot the local magnetic field of my room.",
	},
	{
		Name:        "long",
		Description: "Long Running Command to see how it handles",
	},
}
