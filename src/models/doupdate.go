package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// 一个构件的模型
type Artifact struct {
	Path   string `json:"path"`
	Latest int    `json:"latest"`
	Digest string `json:"digest"`
	Length int64  `json:"length"`
}

// doupdate的主配置文件模型
// 位于.doup/doupdate.json
type DoUpdateConfig struct {
	FingerPrint string     `json:"fingerprint"`
	Latest      int        `json:"latest"`
	Artifacts   []Artifact `json:"artifacts"`
}

var Config DoUpdateConfig

// 从指定文件加载配置文件
func LoadDoUpdateConfig(path string) error {
	jsonb, err := ioutil.ReadFile(path + string(os.PathSeparator) + ".doup" + string(os.PathSeparator) + "doupdate.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonb, &Config)

}

// 将配置信息输出到指定文件
func DumpDoUpdateConfig(path string) error {
	jsonb, err := json.MarshalIndent(GetConfig(), "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path+string(os.PathSeparator)+".doup"+string(os.PathSeparator)+"doupdate.json", jsonb, 0755)
}

func GetConfig() DoUpdateConfig {
	return Config
}
