package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"sensei/internal/commands"
	"sensei/internal/discord"
	"sensei/internal/utils"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	log := logrus.New().WithContext(ctx)
	log = log.WithField("env", utils.GetEnvVar("ENVIRONMENT"))

	// Initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	// Start up discord bot
	token := utils.GetEnvVar("DISCORD_BOT_TOKEN")
	bot := discord.NewBot(token, log)

	ids := utils.GetEnvVar("DEVELOPER_IDS")
	if ids != "" {
		bot.DeveloperIDs = strings.Split(ids, ",")
	}

	defer bot.Stop()

	// Generate Commands
	bot.AddCommands(commands.Ping(log), commands.Pick(log), commands.Roll(log), commands.Join(log), commands.Leave(log))
	bot.RegisterCommands()

	go handleSignalExit()

	// Bind to port
	port := utils.GetEnvVar("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Sensei is online."))
	})
	fmt.Printf("Serving on port %s\n", port)

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

// Wait until CTRL-C or other term signal is received.
func handleSignalExit() {
	fmt.Println("Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	os.Exit(1)
}
