package models

import (
	"doupdate/src/util"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Artifact struct {
	Path    string `json:"path"`
	Latest  int    `json:"latest"`
	Digest  string `json:"digest"`
	Deleted bool   `json:"deleted"`
	Length  int64  `json:"length"`
}

type DoUpdateConfig struct {
	FingerPrint string     `json:"fingerprint"`
	Artifacts   []Artifact `json:"artifacts"`
}

var Config DoUpdateConfig

func LoadDoUpdateConfig(path string) error {
	jsonb, err := ioutil.ReadFile(path + string(os.PathSeparator) + ".doup" + string(os.PathSeparator) + "doupdate.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonb, &Config)

}

func DumpDoUpdateConfig(path string) error {
	jsonb, err := json.MarshalIndent(GetConfig(), "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path+string(os.PathSeparator)+".doup"+string(os.PathSeparator)+"doupdate.json", jsonb, util.OS_USER_RW)
}

func GetConfig() DoUpdateConfig {
	return Config
}
