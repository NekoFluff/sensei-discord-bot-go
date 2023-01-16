package commands

import (
	"fmt"
	"math/rand"
	"sensei/internal/discord"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Pick(log *logrus.Entry) discord.Command {
	command := "pick"
	log = log.WithField("command", command)

	return discord.Command{
		Command: discordgo.ApplicationCommand{
			Name:        command,
			Description: fmt.Sprintf("Pick a random value (e.g. `/%s optionA optionB optionC`)", command),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "options",
					Description: "The several options to choose from",
					Required:    true,
				},
			},
		},
		Handler: func(s discord.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			optionsStr := fmt.Sprint(optionMap["options"].Value)
			args := strings.Split(optionsStr, " ")

			result := pick(args)
			err := respondToInteraction(s, i.Interaction, result)
			if err != nil {
				log.Println(err)
			}
		},
	}
}

func pick(options []string) string {
	return options[rand.Intn(len(options))]
}
