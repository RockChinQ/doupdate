package src

import (
	"doupdate/src/commands"
	"errors"
	"fmt"
)

// 处理命令行程序参数
func Execute(args []string) error {
	if len(args) < 2 { // 未指定操作
		fmt.Println("please specific one of valid commands:")
		for name := range commands.GetCommandList() {
			fmt.Println("- " + name)
		}
		fmt.Println()
		return errors.New("no commands specified")
	}
	// 在已注册的命令列表中查找指令并执行
	for name, cmd := range commands.GetCommandList() {
		if args[1] == name {
			err := cmd(args)
			return err
		}
	}
	return errors.New("no such command: " + args[1])
}
