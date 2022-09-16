package models

type Config struct {
	Bot struct {
		CmdPrefix      string
		VocationEmojis map[string]string `yaml:"vocationEmojis"`
	}
}

type Instance struct {
	Config *Config
}

type Bot struct {
	ID    string
	Token string

	Instances []*Instance
}

type Cmd struct {
	MessageID string
	Message   string
	ChannelID string
	UserID    string
	Command   string
	ArgsRaw   string
	Args      []string
}
