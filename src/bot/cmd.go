package bot

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sergot/tibiago/ent/boss"
	"github.com/sergot/tibiago/src/models"
	"github.com/sergot/tibiago/src/utils"
)

func RunCommand(s *discordgo.Session, cmd *models.Cmd) error {
	// TODO: another package?
	if cmd.Command == "create" {
		client, err := models.ConnectDatabase()
		if err != nil {
			return err
		}
		defer client.Close()

		// if err := client.Schema.Create(context.Background()); err != nil {
		// 	return err
		// }

		// args
		var hours int
		hours, err = strconv.Atoi(cmd.Args[2])
		if err != nil {
			return err
		}

		// TODO: server save time
		hours = 10 + hours

		timestamp, err := time.Parse("2006-01-02 15:04", fmt.Sprintf("%s %d:00", cmd.Args[1], hours))
		if err != nil {
			return err
		}

		b, err := client.Boss.
			Query().
			Where(boss.NameEQ(cmd.Args[0])).
			Only(context.Background())
		if err != nil {
			return err
		}

		bl, err := client.Bosslist.Create().
			SetBoss(b).
			SetStartsAt(timestamp).
			Save(context.Background())
		if err != nil {
			return err
		}
		log.Println("Bosslist created: ", bl)

		list := utils.GenerateBossList(bl)

		bl_msg, err := s.ChannelMessageSend(cmd.ChannelID, list)
		if err != nil {
			return err
		}

		_, err = bl.
			Update().
			SetDiscordMessageID(bl_msg.ID).
			Save(context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}
