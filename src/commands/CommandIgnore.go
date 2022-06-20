package commands

import (
	"doupdate/src/models"
	"errors"
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
	models.Ignored.Rules = append(models.Ignored.Rules, rule[:len(rule)-1])

	return models.DumpIgnoreRules(".")
}
