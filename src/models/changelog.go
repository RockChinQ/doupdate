package models

import (
	"doupdate/src/util"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

type ChangeLog struct {
	Version   int    `json:"version"`
	Changes   string `json:"changes"`
	TimeStamp int64  `json:"timestamp"`
}

func DumpChangeLog(changeLog ChangeLog, path string) error {
	jsonb, err := json.MarshalIndent(changeLog, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path+string(os.PathSeparator)+".doup"+string(os.PathSeparator)+"changelogs"+string(os.PathSeparator)+strconv.Itoa(changeLog.Version)+".json", jsonb, util.OS_USER_RW)
}
