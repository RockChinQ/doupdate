package commands

import (
	"doupdate/src/models"
	"errors"
	"regexp"
)

func CommandIgnore(args []string) error {
	err := models.LoadIgnoreRules(".")
	if err != nil {
		return err
	}

	if len(args) < 3 {
		return errors.New("please provide rule as args[2]")
	}

	rule := ""
	for i := 2; i < len(args); i++ {
		rule += args[i] + " "
	}

	rule = rule[:len(rule)-1]
	//检查是否已存在
	for _, ru := range models.Ignored.Rules {
		if ru == rule {
			return errors.New("rule already exisits:" + rule)
		}
	}

	//检查正则表达式合法性
	_, err = regexp.Compile(rule)
	if err != nil {
		return err
	}

	models.Ignored.Rules = append(models.Ignored.Rules, rule)

	return models.DumpIgnoreRules(".")
}
