package commands

import (
	"fmt"
	"math/rand"
	"regexp"
	"sensei/internal/discord"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Roll(log *logrus.Entry) discord.Command {
	command := "roll"
	log = log.WithField("command", command)

	return discord.Command{
		Command: discordgo.ApplicationCommand{
			Name:        command,
			Description: fmt.Sprintf("Roll some dice (e.g. `/%s 3d20`)", command),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "dice",
					Description: "The dice to roll",
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

			dice := fmt.Sprint(optionMap["dice"].Value)
			r, _ := regexp.Compile(`(?P<count>\d+)d(?P<size>\d+)`)
			matches := r.FindStringSubmatch(dice)

			matchMap := make(map[string]string)

			if len(matches) < 3 {
				err := respondToInteraction(s, i.Interaction, fmt.Sprintf("Invalid roll: %s", dice))
				if err != nil {
					log.Println(err)
				}
				return
			}

			fmt.Println(matches)

			for i, name := range r.SubexpNames() {
				if i != 0 && name != "" {
					matchMap[name] = matches[i]
				}
			}

			count, _ := strconv.Atoi(matchMap["count"])
			size, _ := strconv.Atoi(matchMap["size"])
			result := roll(count, size)
			err := respondToInteraction(s, i.Interaction, fmt.Sprint(result))

			if err != nil {
				log.Println(err)
			}
		},
	}
}

func roll(count int, size int) int {
	if size <= 0 {
		return 0
	}

	total := 0
	for i := 0; i < count; i++ {
		total += rand.Intn(size) + 1
	}
	return total
}
