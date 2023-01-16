package commands

import (
	"sensei/internal/discord"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Ping(log *logrus.Entry) discord.Command {
	command := "ping"
	log = log.WithField("command", command)

	return discord.Command{
		Command: discordgo.ApplicationCommand{
			Name:        command,
			Description: "Is the bot online?",
		},
		Handler: func(s discord.Session, i *discordgo.InteractionCreate) {
			err := respondToInteraction(s, i.Interaction, "Pong!")
			if err != nil {
				log.Println("An error occurred while pinging the server")
				log.Println(err)
			}
		},
	}
}
