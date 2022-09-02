package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	ID    string
	Token string

	CmdPrefix string
}

func Connect() {
	config, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(config)

	token := os.Getenv("BOT_TOKEN")
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating discord session,", err)
		return
	}

	bot := New()

	me, err := s.User("@me")
	if err != nil {
		fmt.Println("Error obtaining account details,", err)
		return
	}
	bot.ID = me.ID

	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuilds | discordgo.IntentGuildMessageReactions)

	s.AddHandler(ready)
	s.AddHandler(messageReactionAdd)

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

func ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Bot is ready")
	fmt.Println(r.User.Username)

	for _, guild := range s.State.Guilds {
		guild_details, err := s.Guild(guild.ID)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("- %s\n", guild_details.Name)
		channels, err := s.GuildChannels(guild.ID)
		if err != nil {
			log.Fatalln(err)
		}

		for _, channel := range channels {
			if channel.Type != discordgo.ChannelTypeGuildText {
				continue
			}
			fmt.Printf("\t- %s\n", channel.Name)
		}
	}
}

func messageReactionAdd(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	r, err := s.ChannelMessageSend(m.ChannelID, m.Emoji.Name)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(r)
}
