package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ignored.json文件的模型
type IgnoreRule struct {
	Rules []string `json:"rules"`
}

var Ignored IgnoreRule

// 从指定文件加载文件忽略规则
func LoadIgnoreRules(path string) error {

	jsonb, err := ioutil.ReadFile(path + string(os.PathSeparator) + ".doup" + string(os.PathSeparator) + "ignored.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonb, &Ignored)
}

// 将现有的文件忽略规则保存到指定文件
func DumpIgnoreRules(path string) error {
	jsonb, err := json.MarshalIndent(Ignored, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path+string(os.PathSeparator)+".doup"+string(os.PathSeparator)+"ignored.json", jsonb, 0755)
}
