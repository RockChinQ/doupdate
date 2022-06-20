package commands

import (
	"doupdate/src/models"
	"doupdate/src/util"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func CommandInitialize(args []string) error {
	targetDir := util.GetStringFromArrayOrDefault(args, 2, ".")

	if !util.Exists(targetDir) {
		return errors.New("no such dir: " + targetDir)
	}

	if !util.IsDir(targetDir) {
		return errors.New("target is not a directory: " + targetDir)
	}

	//开始创建
	err := os.Mkdir(targetDir+string(os.PathSeparator)+".doup", util.OS_USER_RW)
	if err != nil {
		return err
	}

	err = os.Mkdir(targetDir+string(os.PathSeparator)+".doup"+string(os.PathSeparator)+"changelogs", util.OS_USER_RW)
	if err != nil {
		return err
	}

	err = os.Mkdir(targetDir+string(os.PathSeparator)+".doup"+string(os.PathSeparator)+"versions", util.OS_USER_RW)
	if err != nil {
		return err
	}

	//doupdate.json对象
	var doupdate models.DoUpdateConfig

	//生成目录指纹
	doupdate.FingerPrint = util.GenerateFingerPrint()

	doupdate.Artifacts = make([]models.Artifact, 0)

	bytes, err := json.MarshalIndent(doupdate, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(targetDir+string(os.PathSeparator)+".doup"+string(os.PathSeparator)+"doupdate.json", bytes, util.OS_USER_RW)

	if err != nil {
		return err
	}

	var ignored models.IgnoreRule

	ignored.Rules = make([]string, 0)

	bytes, err = json.MarshalIndent(ignored, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(targetDir+string(os.PathSeparator)+".doup"+string(os.PathSeparator)+"ignored.json", bytes, util.OS_USER_RW)

	if err != nil {
		return err
	}

	return nil
}
