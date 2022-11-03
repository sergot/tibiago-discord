package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sergot/tibiago/ent"
	"github.com/sergot/tibiago/src/utils"
)

func UpdateList(s *discordgo.Session, bl *ent.Bosslist, channelID, messageID string) error {
	list := utils.GenerateBosslist(bl)

	_, err := s.ChannelMessageEdit(channelID, messageID, list)
	return err
}
