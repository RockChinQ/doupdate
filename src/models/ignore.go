package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type IgnoreRule struct {
	Rules []string `json:"rules"`
}

var Ignored IgnoreRule

func LoadIgnoreRules(path string) error {

	jsonb, err := ioutil.ReadFile(path + string(os.PathSeparator) + ".doup" + string(os.PathSeparator) + "ignored.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonb, &Ignored)
}

func DumpIgnoreRules(path string) error {
	jsonb, err := json.MarshalIndent(Ignored, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path+string(os.PathSeparator)+".doup"+string(os.PathSeparator)+"ignored.json", jsonb, 0755)
}
