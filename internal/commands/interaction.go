package commands

import (
	"sensei/internal/discord"

	"github.com/bwmarrin/discordgo"
)

func respondToInteraction(s discord.Session, i *discordgo.Interaction, msg string) (err error) {
	return s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
}
