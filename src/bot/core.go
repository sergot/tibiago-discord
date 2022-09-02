package bot

import "os"

func New() *Bot {
	b := &Bot{}

	b.Token = os.Getenv("BOT_TOKEN")
	b.CmdPrefix = "!"

	return b
}

func (b *Bot) LoadConfig() {

}
