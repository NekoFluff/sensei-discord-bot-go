package commands

import (
	"fmt"
	"sensei/internal/discord"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Leave(log *logrus.Entry) discord.Command {
	command := "leave"
	log = log.WithField("command", command)

	return discord.Command{
		Command: discordgo.ApplicationCommand{
			Name:        command,
			Description: fmt.Sprintf("Leave a role (e.g. `/%s @Friends`)", command),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "The role to leave",
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

			hasRole := false
			memberRoleIDs := i.Member.Roles
			for _, mRoleID := range memberRoleIDs {
				if mRoleID == roleID {
					hasRole = true
				}
			}

			if !hasRole {
				err := respondToInteraction(s, i.Interaction, "You don't have that role.")
				if err != nil {
					log.Println(err)
				}
				return
			}

			err := s.GuildMemberRoleRemove(i.GuildID, i.Member.User.ID, roleID)
			if err != nil {
				log.Error("Failed to remove guild member from a role", err)
				err = respondToInteraction(s, i.Interaction, "Sorry, but I don't have the authority to remove you from that role.")
				if err != nil {
					log.Println(err)
				}
				return
			}

			err = respondToInteraction(s, i.Interaction, fmt.Sprintf("You've been removed."))
			if err != nil {
				log.Println(err)
			}
		},
	}
}
