package models

type Config struct {
	Bot struct {
		CmdPrefix string
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
