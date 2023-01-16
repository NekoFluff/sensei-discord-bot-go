package commands

import (
	"context"
	"fmt"
	"sensei/internal/discord"
	"testing"

	"github.com/bwmarrin/discordgo"
	gomock "github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
)

func TestServer_Ping(t *testing.T) {
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

		log := logrus.New().WithContext(context.TODO())
		session := discord.NewMockSession(ctrl)
		tt.setupMock(session)

		ping := Ping(log)
		ping.Handler(session, &discordgo.InteractionCreate{
			Interaction: &discordgo.Interaction{
				ChannelID: "test-channel-id",
			},
		})
	}
}
