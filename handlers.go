package main

import (
	"testgo/cmd"

	"github.com/bwmarrin/discordgo"
)

var all_handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"mensa":                    cmd.MensaHander,
	"terminal":                 cmd.TerminalHander,
	"impressum":                cmd.ImpressumHandler,
	//"basic-command-with-files": cmd.BasicFileHander,
	"b-field":                  cmd.BFieldHandler,
	"long":                     cmd.LongRunningHandler,
	"ozon":						cmd.Ozon,
}
