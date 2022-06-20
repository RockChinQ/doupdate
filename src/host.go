package src

import (
	"doupdate/src/commands"
	"errors"
	"fmt"
)

func Execute(args []string) error {
	if len(args) < 2 {
		fmt.Println("please specific one of valid commands:")
		for name := range commands.GetCommandList() {
			fmt.Println("- " + name)
		}
		fmt.Println()
		return errors.New("no commands specified")
	}

	for name, cmd := range commands.GetCommandList() {
		if args[1] == name {
			err := cmd(args)
			return err
		}
	}
	return errors.New("no such command: " + args[1])
}
