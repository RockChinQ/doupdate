package commands

import "fmt"

var help = "\nusage:\n" +
	"- init           \tInitialize repository\n" +
	"- ignore <regexp>\tAdd a path ignoring rule(RegExp)\n" +
	"- status        \tSee repository changes\n" +
	"- release        \tPublish a new version\n" +
	"- log            \tSee version change logs\n" +
	"- help           \tPrint this message\n"

func CommandHelp(args []string) error {

	fmt.Println(help)

	return nil
}
