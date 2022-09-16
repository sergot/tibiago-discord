package bot

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/sergot/tibiago/ent/bosslist"
	"github.com/sergot/tibiago/ent/participant"
	"github.com/sergot/tibiago/src/models"
	"github.com/sergot/tibiago/src/utils"
)

func ReadyHandler(db string) func(s *discordgo.Session, r *discordgo.Ready) {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
		fmt.Println(r.User.Username)

		for _, guild := range s.State.Guilds {
			guild_details, err := s.Guild(guild.ID)
			if err != nil {
				log.Println(err)
			}

			instances_map[guild.ID] = &models.Instance{
				Config: defaultConfig,
			}

			fmt.Printf("- %s\n", guild_details.Name)

			channels, err := s.GuildChannels(guild.ID)
			if err != nil {
				log.Println(err)
				return
			}

			for _, channel := range channels {
				if channel.Type != discordgo.ChannelTypeGuildText {
					continue
				}
				fmt.Printf("\t- %s\n", channel.Name)
			}
		}
	}
}

func ReactionRemoveHandler(db string) func(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
	return func(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
		message, err := s.ChannelMessage(m.ChannelID, m.MessageID)
		if err != nil {
			log.Println(err)
		}

		me, err := s.User("@me")
		if err != nil {
			log.Println("Error obtaining account details,", err)
		}

		if me.ID != message.Author.ID {
			return
		}

		client, err := models.ConnectDatabase()
		if err != nil {
			log.Println(err)
		}
		defer client.Close()

		bl, err := client.Bosslist.
			Query().
			Where(bosslist.DiscordMessageID(m.MessageID)).
			WithParticipants().
			Only(context.Background())
		if err != nil {
			log.Println(err)
		}

		p, err := bl.QueryParticipants().
			Where(participant.DiscordID(m.UserID)).
			// Where(vocation)
			Only(context.Background())
		if err != nil {
			log.Println(err)
		}

		err = client.Participant.
			DeleteOneID(p.ID).
			Exec(context.Background())
		if err != nil {
			log.Println(err)
		}

		list := utils.GenerateBosslist(bl)

		_, err = s.ChannelMessageEdit(message.ChannelID, message.ID, list)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func ReactionAddHandler(db string) func(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	return func(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
		message, err := s.ChannelMessage(m.ChannelID, m.MessageID)
		if err != nil {
			log.Println(err)
		}

		me, err := s.User("@me")
		if err != nil {
			log.Println("Error obtaining account details,", err)
		}

		if me.ID != message.Author.ID {
			return
		}

		instance := instances_map[m.GuildID]
		vocation := instance.Config.Bot.VocationEmojis[m.Emoji.Name]
		fmt.Println(instance.Config.Bot.VocationEmojis)
		if vocation == "" {
			log.Println("Unknown vocation emoji: ", m.Emoji.APIName())
			return
		}

		client, err := models.ConnectDatabase()
		if err != nil {
			log.Println(err)
		}
		defer client.Close()

		p, err := client.Participant.
			Create().
			SetVocation(participant.Vocation(vocation)).
			SetDiscordID(m.UserID).
			Save(context.Background())
		if err != nil {
			_ = s.MessageReactionRemove(m.ChannelID, m.MessageID, m.Emoji.ID, m.UserID)
			// TODO: handle unique contraint violation (priv msg to the user?)
			log.Println(err)
			return
		}

		bl, err := client.Bosslist.
			Query().
			Where(bosslist.DiscordMessageID(m.MessageID)).
			Only(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		_, err = bl.Update().
			AddParticipants(p).
			Save(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		list := utils.GenerateBosslist(bl)
		fmt.Println(list)

		_, err = s.ChannelMessageEdit(message.ChannelID, message.ID, list)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func CommonHandler(db string) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		instance := instances_map[m.GuildID]
		content := m.Content

		fmt.Println(content)

		cmd, err := utils.ParseCmd(m, instance)
		if err != nil {
			log.Println(err)
			return
		}

		err = RunCommand(s, cmd)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
