package models

type Cmd struct {
	MessageID string
	Message   string
	ChannelID string
	UserID    string
	Command   string
	ArgsRaw   string
	Args      []string
}
