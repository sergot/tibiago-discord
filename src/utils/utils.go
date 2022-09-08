package utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/bwmarrin/discordgo"
	"github.com/mattn/go-shellwords"
	"github.com/mozillazg/go-unidecode"

	"github.com/sergot/tibiago/ent"
	"github.com/sergot/tibiago/src/models"
)

func ParseCmd(m *discordgo.MessageCreate, instance *models.Instance) (*models.Cmd, error) {
	content := m.Message.Content

	if !strings.HasPrefix(content, instance.Config.Bot.CmdPrefix) {
		return nil, nil
	}

	s := strings.TrimSpace(content)

	c := &models.Cmd{
		MessageID: m.Message.ID,
		Message:   strings.TrimSpace(strings.TrimPrefix(s, instance.Config.Bot.CmdPrefix)),
		ChannelID: strings.TrimSpace(m.ChannelID),
		UserID:    m.Author.ID,
	}

	if c.Message == "" {
		return nil, nil
	}

	firstOccurrence := true
	firstUnicodeSpace := func(c rune) bool {
		isFirstSpace := unicode.IsSpace(c) && firstOccurrence
		if isFirstSpace {
			firstOccurrence = false
		}
		return isFirstSpace
	}

	pieces := strings.FieldsFunc(c.Message, firstUnicodeSpace)
	c.Command = strings.ToLower(unidecode.Unidecode(pieces[0]))

	if len(pieces) > 1 {
		c.ArgsRaw = strings.TrimSpace(pieces[1])
		parsedArgs, err := shellwords.Parse(c.ArgsRaw)
		if err != nil {
			return nil, errors.New("error parsing arguments: " + err.Error())
		}
		c.Args = parsedArgs
	}
	return c, nil
}

func PrintParticipants(template string, participants []*ent.Participant) []string {
	vocs := regexp.MustCompile(`[a-zA-Z]+`).FindAllString(template, -1)
	amounts := regexp.MustCompile(`[0-9]+`).FindAllString(template, -1)

	var result []string
	for i, voc := range vocs {
		n, err := strconv.Atoi(amounts[i])
		if err != nil {
			return nil
		}
		for j := 0; j < n; j++ {
			user := voc
			for x, p := range participants {
				if p.Vocation.String() == voc {
					user = fmt.Sprintf("<@%s>", p.DiscordID)
					participants = append(participants[:x], participants[x+1:]...)
					break
				}
			}

			result = append(result, user)
		}
	}

	return result
}

func GenerateBossList(bosslist *ent.Bosslist) string {
	boss := bosslist.QueryBoss().OnlyX(context.Background())

	t_data := struct {
		Name      string
		Timestamp string
		// TODO: is this secure?
		Participants []template.HTML
	}{
		Name:      boss.Name,
		Timestamp: bosslist.StartsAt.Format("2006-01-02 15:04"),
	}

	tmpl := bosslist.CustomTemplate
	if tmpl == "" {
		tmpl = boss.Template
	}

	participants := PrintParticipants(tmpl, bosslist.QueryParticipants().AllX(context.Background()))
	fmt.Println("PATO:", participants)

	for _, p := range participants {
		t_data.Participants = append(t_data.Participants, template.HTML(p))
	}

	buf := new(bytes.Buffer)
	t, err := template.ParseFiles("templates/bosslist.tmpl")
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(buf, t_data)
	if err != nil {
		log.Println(err)
	}

	return buf.String()
}
