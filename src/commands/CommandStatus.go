package commands

import (
	"doupdate/src/models"
	"doupdate/src/util"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func CommandStatus(args []string) error {
	err := models.LoadDoUpdateConfig(".")
	if err != nil {
		return err
	}

	err = models.LoadIgnoreRules(".")
	if err != nil {
		return err
	}

	newList, updatedList, deletedList, err := GetStatusList()
	if err != nil {
		return err
	}

	if len(newList) != 0 {
		fmt.Println("\nadded files:")
		for n, filePath := range newList {
			fmt.Println(strconv.Itoa(n+1) + " " + filePath)
		}
	}

	if len(updatedList) != 0 {
		fmt.Println("\nupdated files:")
		for n, filePath := range updatedList {
			fmt.Println(strconv.Itoa(n+1) + " " + filePath)
		}
	}

	if len(deletedList) != 0 {
		fmt.Println("\ndeleted files:")
		for n, filePath := range deletedList {
			fmt.Println(strconv.Itoa(n+1) + " " + filePath)
		}
	}
	if len(newList) == 0 && len(updatedList) == 0 && len(deletedList) == 0 {
		fmt.Println("no file changed")
	} else {
		fmt.Println("")
	}

	return nil
}

//新增,更新,删除
func GetStatusList() ([]string, []string, []string, error) {

	fileList := make([]string, 0)

	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {

		if strings.Contains(path, ".doup") {
			return nil
		}

		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fileList = append(fileList, path)
		return nil
	})

	if err != nil {
		return nil, nil, nil, err
	}

	newList := make([]string, 0)
	updatedList := make([]string, 0)
	deletedList := make([]string, 0)

	//查找与配置内的差异
	cfg := models.GetConfig()

	//遍历文件列表
nextFile:
	for _, filePath := range fileList {
		//遍历构件列表
		for _, artifact := range cfg.Artifacts {
			if artifact.Path == filePath {
				//计算md5
				hash, err := util.FileMD5(filePath)
				if err != nil {
					return nil, nil, nil, err
				}

				if hash != artifact.Digest {
					updatedList = append(updatedList, filePath)
				}
				continue nextFile
			}
		}
		//新文件

		//检查是不是在忽略列表
		for _, rule := range models.Ignored.Rules {
			ignored, err := PathMatch(rule, filePath)
			if err != nil {
				return nil, nil, nil, err
			}
			if ignored {
				continue nextFile
			}
		}
		newList = append(newList, filePath)
	}

	//遍历构件列表,检查已删除的
nextArtifact:
	for _, artifact := range cfg.Artifacts {
		for _, filePath := range fileList {
			if artifact.Path == filePath {
				continue nextArtifact
			}
		}
		deletedList = append(deletedList, artifact.Path)
	}

	return newList, updatedList, deletedList, nil
}

func PathMatch(ruleReg, path string) (bool, error) {
	return regexp.MatchString(ruleReg, path)
}
