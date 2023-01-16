package commands

import (
	"fmt"
	"sensei/internal/discord"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Join(log *logrus.Entry) discord.Command {
	command := "join"
	log = log.WithField("command", command)

	return discord.Command{
		Command: discordgo.ApplicationCommand{
			Name:        command,
			Description: fmt.Sprintf("Join a role (e.g. `/%s @Friends`)", command),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "The role to join",
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

			roleID := fmt.Sprint(optionMap["role"].Value)

			memberRoleIDs := i.Member.Roles
			for _, mRoleID := range memberRoleIDs {
				if mRoleID == roleID {
					err := respondToInteraction(s, i.Interaction, "You already have that role.")
					if err != nil {
						log.Println(err)
					}
					return
				}
			}

			err := s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, roleID)
			if err != nil {
				log.Error("Failed to add guild member to a role", err)
				err = respondToInteraction(s, i.Interaction, "Sorry, but I don't have the authority to add you to that role.")
				if err != nil {
					log.Println(err)
				}
				return
			}

			err = respondToInteraction(s, i.Interaction, "You've been added.")
			if err != nil {
				log.Println(err)
			}
		},
	}
}
