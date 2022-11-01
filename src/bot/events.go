package bot

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/sergot/tibiago/ent"
	"github.com/sergot/tibiago/ent/bosslist"
	"github.com/sergot/tibiago/ent/instance"
	"github.com/sergot/tibiago/ent/participant"
	"github.com/sergot/tibiago/src/models"
	"github.com/sergot/tibiago/src/utils"
)

func initGuild(client *ent.Client, guildID string) error {
	instance, _ := client.Instance.
		Query().
		Where(instance.DiscordGuildID(guildID)).
		WithConfigs().
		Only(context.Background())

	if instance == nil {
		log.Println("Instance not found, initializing new one")

		tmp := client.Instance.
			Create().
			SetDiscordGuildID(guildID)

		instance = tmp.SaveX(context.Background())

		defaultConfig, err := LoadConfig("default_config.yaml")
		if err != nil {
			return err
		}
		configs := utils.MapConfigToDBConfig(defaultConfig)
		bulk := make([]*ent.InstanceConfigCreate, len(configs))
		for i, c := range configs {
			bulk[i] = client.InstanceConfig.Create().
				SetInstance(instance).
				SetKey(c.Key).
				SetValue(c.Value)
		}

		_, err = client.InstanceConfig.CreateBulk(bulk...).Save(context.Background())
		if err != nil {
			return err
		}
	}

	log.Println("Making the instance active")
	instance.Update().
		SetStatus("active").
		SaveX(context.Background())

	return nil
}

func GuildCreateHandler(db string) func(s *discordgo.Session, e *discordgo.GuildCreate) {
	return func(_ *discordgo.Session, e *discordgo.GuildCreate) {
		client, err := models.ConnectDatabase()
		if err != nil {
			log.Println(err)
		}

		err = initGuild(client, e.ID)
		if err != nil {
			log.Println(err)
		}

		fmt.Println("Joined new guild")
		fmt.Println(e.Guild.Name)
	}
}

func ReadyHandler(db string) func(s *discordgo.Session, r *discordgo.Ready) {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
		fmt.Println(r.User.Username)

		for _, guild := range s.State.Guilds {
			guild_details, err := s.Guild(guild.ID)
			if err != nil {
				log.Println(err)
			}

			// client, err := models.ConnectDatabase()
			// if err != nil {
			// 	log.Println(err)
			// }

			// instances_map[guild.ID] = &models.Instance{
			// 	Config: defaultConfig,
			// }

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

		client, err := models.ConnectDatabase()
		if err != nil {
			log.Println(err)
		}
		defer client.Close()

		instance := client.Instance.
			Query().
			Where(instance.DiscordGuildID(m.GuildID)).
			OnlyX(context.Background())

		config := utils.MapDBConfigToConfig(instance.QueryConfigs().AllX(context.Background()))

		vocation := config.Bot.VocationEmojis[m.Emoji.Name]
		if vocation == "" {
			log.Println("Unknown vocation emoji: ", m.Emoji.APIName())
			return
		}

		p, err := client.Participant.
			Create().
			SetVocation(participant.Vocation(vocation)).
			SetDiscordID(m.UserID).
			Save(context.Background())
		if err != nil {
			_ = s.MessageReactionRemove(m.ChannelID, m.MessageID, m.Emoji.APIName(), m.UserID)
			log.Println("Same person reacts with the same emoji", err)
			return
		}

		bl, err := client.Bosslist.
			Query().
			Where(bosslist.DiscordMessageID(m.MessageID)).
			WithParticipants().
			Only(context.Background())
		if err != nil {
			log.Println("Bosslist not found", err)
			return
		}

		_, err = bl.Update().
			AddParticipants(p).
			Save(context.Background())
		if err != nil {
			log.Println("Failed to add participant:", err)
			err = client.Participant.DeleteOne(p).Exec(context.Background())
			if err != nil {
				log.Println("Failed to delete participant:", err)
			}
			_ = s.MessageReactionRemove(m.ChannelID, m.MessageID, m.Emoji.APIName(), m.UserID)
		}

		list := utils.GenerateBosslist(bl)

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

		client, err := models.ConnectDatabase()
		if err != nil {
			log.Println(err)
		}
		defer client.Close()

		instance := client.Instance.
			Query().
			Where(instance.DiscordGuildID(m.GuildID)).
			OnlyX(context.Background())

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
