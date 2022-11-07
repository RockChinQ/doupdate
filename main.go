package main

import (
	"doupdate/src"
	"os"
)

// 程序入口
// 将命令行参数输入到命令处理器
func main() {
	err := src.Execute(os.Args)
	if err != nil {
		panic(err)
	}
}
