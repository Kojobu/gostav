package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "YOUR TOKEN", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {

	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand{
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
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){

		"mensa": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: mensa_scrap(false),
				},
			})
		},

		"terminal": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

		},

		"test-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: test(),
				},
			})
		},

		"impressum": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: impressum(),
				},
			})
		},

		"basic-command-with-files": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
		},

		"b-field": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
		},
	}
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		log.Println("Removing commands...")
		// // We need to fetch the commands, since deleting requires the command ID.
		// // We are doing this from the returned commands on line 375, because using
		// // this will delete all the commands, which might not be desirable, so we
		// // are deleting only the commands that we added.
		// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
		// if err != nil {
		// 	log.Fatalf("Could not fetch registered commands: %v", err)
		// }

		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Goodbye Gostav. :)")
}
