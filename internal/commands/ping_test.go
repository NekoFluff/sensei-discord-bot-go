package commands

import (
	"fmt"
	"os"
	"sensei/internal/discord"
	"testing"

	"github.com/bwmarrin/discordgo"
	gomock "github.com/golang/mock/gomock"
)

func TestServer_Ping(t *testing.T) {
	os.Setenv("COMMAND_PREFIX", "!")

	tests := []struct {
		name      string
		setupMock func(*discord.MockSession)
	}{
		{
			name: "successfully pinged server",
			setupMock: func(session *discord.MockSession) {
				session.EXPECT().InteractionRespond(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
		},
		{
			name: "failed to ping server",
			setupMock: func(session *discord.MockSession) {
				session.EXPECT().InteractionRespond(gomock.Any(), gomock.Any()).Times(1).Return(fmt.Errorf("random error"))
			},
		},
	}

	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		session := discord.NewMockSession(ctrl)
		tt.setupMock(session)

		ping := Ping(nil)
		ping.Handler(session, &discordgo.InteractionCreate{
			Interaction: &discordgo.Interaction{
				ChannelID: "test-channel-id",
			},
		})
	}
}
