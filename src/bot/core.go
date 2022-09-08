package bot

import (
	"os"

	"github.com/sergot/tibiago/src/models"
)

func New() *models.Bot {
	b := &models.Bot{}

	b.Token = os.Getenv("BOT_TOKEN")

	return b
}
