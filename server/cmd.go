package server

type commandID int

const (
	cmdName commandID = iota
	cmdJoin
	cmdList
	cmdMsg
	cmdQuit
	cmdHelp
)

type cmd struct {
	id         commandID
	connection *connection
	args       []string
}
