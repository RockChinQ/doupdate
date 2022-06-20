package commands

type Command func(args []string) error

var commands map[string]Command

//把所有可用的指令注册
func init() {
	commands = make(map[string]Command)

	commands["init"] = CommandInitialize
	commands["status"] = CommandStatus
	commands["release"] = CommandRelease
	commands["ignore"] = CommandIgnore
	commands["log"] = CommandLog
}

func GetCommandList() map[string]Command {
	return commands
}
