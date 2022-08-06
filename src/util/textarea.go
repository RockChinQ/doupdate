package util

import (
	"fmt"
	"os"
)

func GetText(tips string) string {
	buffer := ""

	fmt.Println(tips)
	fmt.Println("(type 'f' in a line to finish,while 'e' to exit)")
	for {
		line := ""
		_, _ = fmt.Scanln(&line)

		if line == "f" {
			if len(buffer) == 0 {
				return ""
			}
			return buffer[:len(buffer)-1]
		} else if line == "e" {
			os.Exit(0)
		}

		buffer += line + "\n"
	}
}
