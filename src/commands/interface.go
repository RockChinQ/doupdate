package commands

// 定义一个指令的函数模型
type Command func(args []string) error

// 指令列表
var commands map[string]Command

//把所有可用的指令注册
func init() {
	commands = make(map[string]Command)

	commands["init"] = CommandInitialize
	commands["status"] = CommandStatus
	commands["release"] = CommandRelease
	commands["ignore"] = CommandIgnore
	commands["log"] = CommandLog
	commands["help"] = CommandHelp
}

func GetCommandList() map[string]Command {
	return commands
}
