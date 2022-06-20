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
	maxVersion := 0
	for _, artifact := range models.GetConfig().Artifacts {
		if artifact.Latest > maxVersion {
			maxVersion = artifact.Latest
		}
	}

	versionNum := 0

	if util.GetNowDateTime() == maxVersion/100 {
		versionNum = util.GetNowDateTime()*100 + (maxVersion%100 + 1)
	} else {
		versionNum = util.GetNowDateTime() * 100
	}

	//获取changelog

	changelog := models.ChangeLog{
		Version:   versionNum,
		Changes:   util.GetText("enter changelogs for version " + strconv.Itoa(versionNum)),
		TimeStamp: time.Now().Unix(),
	}
	err = models.DumpChangeLog(changelog, ".")
	if err != nil {
		return err
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
			Path:    filePath,
			Digest:  digest,
			Latest:  versionNum,
			Deleted: false,
			Length:  length,
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
				models.Config.Artifacts[n].Deleted = false
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

	//标记已删除的
	for _, filePath := range deletedList {
		for n, artifact := range models.Config.Artifacts {
			if artifact.Path == filePath {
				models.Config.Artifacts[n].Deleted = true
				models.Config.Artifacts[n].Latest = versionNum
				break
			}
		}
	}

	return models.DumpDoUpdateConfig(".")
}
