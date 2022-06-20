package main

import (
	"doupdate/src"
	"os"
)

func main() {
	err := src.Execute(os.Args)
	if err != nil {
		panic(err)
	}
}
