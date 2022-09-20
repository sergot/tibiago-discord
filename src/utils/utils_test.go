package utils_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sergot/tibiago/carrier"
	"github.com/sergot/tibiago/ent"
	"github.com/sergot/tibiago/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestBosslistParticipants(t *testing.T) {
	type testCase struct {
		template     string
		participants []utils.Part
		expected     []string
	}

	var testCases = []testCase{
		{
			template:     "1ek1ed3shooter",
			participants: []utils.Part{{ID: "1", Voc: "ed"}},
			expected:     []string{"ek", "<@1>", "shooter", "shooter", "shooter"},
		},

		{
			template:     "1ek1ed3shooter",
			participants: []utils.Part{{ID: "1", Voc: "ms"}},
			expected:     []string{"ek", "ed", "<@1>", "shooter", "shooter"},
		},

		{
			template:     "1ek1ed3any",
			participants: []utils.Part{{ID: "1", Voc: "ek"}, {ID: "2", Voc: "ek"}},
			expected:     []string{"<@1>", "ed", "<@2>", "any", "any"},
		},

		{
			template:     "1ek1ed3shooter",
			participants: []utils.Part{{ID: "1", Voc: "ek"}, {ID: "2", Voc: "ek"}},
			expected:     []string{"<@1>", "ed", "shooter", "shooter", "shooter"},
		},

		{
			template:     "1ek1ed3any",
			participants: []utils.Part{{ID: "1", Voc: "ek"}, {ID: "2", Voc: "rp"}},
			expected:     []string{"<@1>", "ed", "<@2>", "any", "any"},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			utils.BosslistParticipants(tc.template, tc.participants),
			tc.expected,
		)
	}
}

func TestGenerateBosslist(t *testing.T) {
	factory := initFactory()

	b, err := factory.BossFactory().SetName("feru").Create(context.TODO())
	if err != nil {
		t.Fail()
	}

	bl, err := factory.BosslistFactory().SetBoss(b).Create(context.TODO())
	if err != nil {
		t.Fail()
	}

	p, err := factory.ParticipantFactory().
		SetDiscordID("1").
		SetVocation("ek").
		Create(context.TODO())
	if err != nil {
		t.Fail()
	}

	bl.Update().AddParticipants(p)

	r := utils.GenerateBosslist(bl)
	assert.Equal(t, r, "feru at 2022-09-09 00:00\n\n\nek\n\ned\n\nshooter\n\nshooter\n\nshooter\n")
}

func initFactory() *carrier.EntFactory {
	client, err := ent.Open("sqlite3", ":memory:?_fk=1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	factory := carrier.NewEntFactory(client)
	initBossFactory(factory)
	initBosslistFactory(factory)
	initParticipantFactory(factory)
	initInstanceFactory(factory)
	initInstanceConfigFactory(factory)
	return factory
}

func initBossFactory(f *carrier.EntFactory) {
	meta := carrier.EntBossMetaFactory().
		SetNameDefault("carrier")
	f.SetBossFactory(meta.Build())
}

func initBosslistFactory(f *carrier.EntFactory) {
	startsat, _ := time.Parse("2006-01-02", "2022-09-09")
	meta := carrier.EntBosslistMetaFactory().
		SetStartsAtDefault(startsat)
	f.SetBosslistFactory(meta.Build())
}

func initParticipantFactory(f *carrier.EntFactory) {
	meta := carrier.EntParticipantMetaFactory().
		SetDiscordIDDefault("1").
		SetVocationDefault("ek")
	f.SetParticipantFactory(meta.Build())
}

func initInstanceFactory(f *carrier.EntFactory) {
	meta := carrier.EntInstanceMetaFactory().
		SetDiscordGuildIDDefault("1").
		SetStatusDefault("active")
	f.SetInstanceFactory(meta.Build())
}

func initInstanceConfigFactory(f *carrier.EntFactory) {
	meta := carrier.EntInstanceConfigMetaFactory()
	f.SetInstanceConfigFactory(meta.Build())
}

func TestParseCmd(t *testing.T) {
	factory := initFactory()

	instance, err := factory.InstanceFactory().
		Create(context.TODO())
	if err != nil {
		t.Fail()
	}

	_, err = factory.InstanceConfigFactory().
		SetKey("bot.cmdprefix").
		SetValue("!").
		SetInstance(instance).
		Create(context.TODO())
	if err != nil {
		t.Fail()
	}

	messageCreate := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "t",
			Content:   "!create a b c d e",
			ChannelID: "t",
			Author: &discordgo.User{
				ID: "t",
			},
		},
	}

	cmd, err := utils.ParseCmd(messageCreate, instance)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Command, "create")
	assert.Equal(t, len(cmd.Args), 5)
}
