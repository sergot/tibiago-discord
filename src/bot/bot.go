package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func Connect(db string) {
	var err error

	token := os.Getenv("BOT_TOKEN")
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating discord session,", err)
		return
	}

	bot := NewBot()

	me, err := s.User("@me")
	if err != nil {
		fmt.Println("Error obtaining account details,", err)
		return
	}
	bot.ID = me.ID

	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuilds | discordgo.IntentGuildMessages | discordgo.IntentGuildMessageReactions)

	s.AddHandler(GuildCreateHandler(db))
	s.AddHandler(ReadyHandler(db))

	s.AddHandler(ReactionAddHandler(db))
	s.AddHandler(CommonHandler(db))

	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}
