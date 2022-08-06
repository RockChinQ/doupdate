package commands

import (
	"doupdate/src/models"
	"doupdate/src/util"
	"fmt"
	"os"
	"strconv"
	"time"
)

func CommandRelease(args []string) error {
	err := models.LoadDoUpdateConfig(".")
	if err != nil {
		return err
	}

	//加载忽略路径的规则
	err = models.LoadIgnoreRules(".")
	if err != nil {
		return err
	}

	newList, updatedList, deletedList, err := GetStatusList()
	if err != nil {
		return err
	}

	if len(newList) == 0 && len(updatedList) == 0 && len(deletedList) == 0 {
		fmt.Println("no file changed")
		return nil
	}

	//生成版本号
	maxVersion := models.Config.Latest

	versionNum := 0

	if util.GetNowDateTime() == maxVersion/100 {
		versionNum = util.GetNowDateTime()*100 + (maxVersion%100 + 1)
	} else {
		versionNum = util.GetNowDateTime() * 100
	}

	models.Config.Latest = versionNum

	//输入changelog
	changelog := models.ChangeLog{
		Version:   versionNum,
		Previous:  maxVersion,
		Changes:   util.GetText("enter changelogs for version " + strconv.Itoa(versionNum)),
		TimeStamp: time.Now().Unix(),
		Brief:     make(map[string][]string),
	}

	//把所有新增和改动的文件存进versions
	if !util.Exists(".doup" + string(os.PathSeparator) + "versions" + string(os.PathSeparator) + strconv.Itoa(versionNum)) {
		err = os.Mkdir(".doup"+string(os.PathSeparator)+"versions"+string(os.PathSeparator)+strconv.Itoa(versionNum), util.OS_USER_RW)
		if err != nil {
			return err
		}
	}

	//添加新增的
	for _, filePath := range newList {
		digest, err := util.FileMD5(filePath)
		if err != nil {
			return err
		}

		length, err := util.GetFileSize(filePath)
		if err != nil {
			return err
		}

		artifact := models.Artifact{
			Path:   filePath,
			Digest: digest,
			Latest: versionNum,
			Length: length,
		}

		//保存到版本库

		_, err = util.Copy(filePath, ".doup"+string(os.PathSeparator)+"versions"+string(os.PathSeparator)+strconv.Itoa(versionNum)+string(os.PathSeparator)+filePath)
		if err != nil {
			return err
		}

		models.Config.Artifacts = append(models.GetConfig().Artifacts, artifact)
	}

	//修改已更新的
	for _, filePath := range updatedList {
		digest, err := util.FileMD5(filePath)
		if err != nil {
			return err
		}
		length, err := util.GetFileSize(filePath)
		if err != nil {
			return err
		}

		for n, artifact := range models.Config.Artifacts {
			if artifact.Path == filePath {
				models.Config.Artifacts[n].Digest = digest
				models.Config.Artifacts[n].Latest = versionNum
				models.Config.Artifacts[n].Length = length
				break
			}
		}

		//保存到版本库

		_, err = util.Copy(filePath, ".doup"+string(os.PathSeparator)+"versions"+string(os.PathSeparator)+strconv.Itoa(versionNum)+string(os.PathSeparator)+filePath)
		if err != nil {
			return err
		}

	}

	//删除已删除的文件的记录
	for _, filePath := range deletedList {
		for n := 0; n < len(models.Config.Artifacts); {
			if models.Config.Artifacts[n].Path == filePath {
				models.Config.Artifacts = append(models.Config.Artifacts[:n], models.Config.Artifacts[n+1:]...)
				break
			} else {
				n++
			}
		}
	}

	changelog.Brief["added"] = newList
	changelog.Brief["updated"] = updatedList
	changelog.Brief["deleted"] = deletedList

	err = models.DumpChangeLog(changelog, ".")
	if err != nil {
		return err
	}
	return models.DumpDoUpdateConfig(".")
}
