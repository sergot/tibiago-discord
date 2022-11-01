package utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"log"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/exp/slices"

	"github.com/bwmarrin/discordgo"
	"github.com/mattn/go-shellwords"
	"github.com/mozillazg/go-unidecode"
	"github.com/sergot/tibiago/ent"
	"github.com/sergot/tibiago/src/models"
	"github.com/thoas/go-funk"
)

func ParseCmd(m *discordgo.MessageCreate, instance *ent.Instance) (*models.Cmd, error) {
	content := m.Message.Content

	config := MapDBConfigToConfig(instance.QueryConfigs().AllX(context.Background()))

	if !strings.HasPrefix(content, config.Bot.CmdPrefix) {
		return nil, nil
	}

	s := strings.TrimSpace(content)

	c := &models.Cmd{
		MessageID: m.Message.ID,
		Message:   strings.TrimSpace(strings.TrimPrefix(s, config.Bot.CmdPrefix)),
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

type Part struct {
	ID  string
	Voc string
}

var aliases map[string][]string = map[string][]string{
	"shooter": {"ed", "ms", "rp"},
	"any":     {"ek", "ed", "ms", "rp"},
}

func BosslistParticipants(template string, parts []Part) []string {
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
			for x, p := range parts {
				if p.Voc == voc || slices.Contains(aliases[voc], p.Voc) {
					user = fmt.Sprintf("<@%s>", p.ID)
					parts = append(parts[:x], parts[x+1:]...)
					break
				}
			}

			result = append(result, user)
		}
	}

	return result
}

func GenerateBosslist(bosslist *ent.Bosslist) string {
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

	all_participants := bosslist.QueryParticipants().AllX(context.Background())
	parts := funk.Map(all_participants, func(ep *ent.Participant) Part {
		return Part{
			ID:  ep.DiscordID,
			Voc: ep.Vocation.String(),
		}
	})
	participants := BosslistParticipants(tmpl, parts.([]Part))
	fmt.Println("PATO:", participants)

	for _, p := range participants {
		t_data.Participants = append(t_data.Participants, template.HTML(p))
	}

	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)

	buf := new(bytes.Buffer)
	// TODO: fix path
	t, err := template.ParseFiles(fmt.Sprintf("%s/../../templates/bosslist.tmpl", basepath))
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(buf, t_data)
	if err != nil {
		log.Println(err)
	}

	return buf.String()
}

// TODO: think whether we can write it better. More dynamic?
func MapConfigToDBConfig(config *models.Config) []*ent.InstanceConfig {

	var result []*ent.InstanceConfig

	result = append(result, &ent.InstanceConfig{
		Key:   "bot.cmdprefix",
		Value: config.Bot.CmdPrefix,
	})

	for k, v := range config.Bot.VocationEmojis {
		result = append(result, &ent.InstanceConfig{
			Key:   fmt.Sprintf("bot.vocationemojis.%s", k),
			Value: v,
		})
	}

	return result
}

func MapDBConfigToConfig(configs []*ent.InstanceConfig) *models.Config {
	result := &models.Config{}
	result.Bot.VocationEmojis = make(map[string]string)

	for _, c := range configs {
		switch {
		case c.Key == "bot.cmdprefix":
			result.Bot.CmdPrefix = c.Value
		case strings.HasPrefix(c.Key, "bot.vocationemojis."):
			if strings.HasPrefix(c.Key, "bot.vocationemojis.") {
				voc := strings.TrimPrefix(c.Key, "bot.vocationemojis.")
				fmt.Println(voc)
				result.Bot.VocationEmojis[voc] = c.Value
			}
		}
	}

	return result
}
