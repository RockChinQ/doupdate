package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

// 版本日志的模型
type ChangeLog struct {
	Version   int                   `json:"version"`
	Previous  int                   `json:"previous"`
	Changes   string                `json:"changes"`
	TimeStamp int64                 `json:"timestamp"`
	Brief     map[string]([]string) `json:"brief"`
}

// 将一个版本日志保存到指定文件
func DumpChangeLog(changeLog ChangeLog, path string) error {
	jsonb, err := json.MarshalIndent(changeLog, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path+string(os.PathSeparator)+".doup"+string(os.PathSeparator)+"changelogs"+string(os.PathSeparator)+strconv.Itoa(changeLog.Version)+".json", jsonb, 0755)
}
