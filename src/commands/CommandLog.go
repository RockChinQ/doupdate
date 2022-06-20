package commands

import (
	"doupdate/src/models"
	"doupdate/src/util"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func CommandLog(args []string) error {

	versions := make([]int, 0)

	err := filepath.Walk(".doup"+string(os.PathSeparator)+"changelogs"+string(os.PathSeparator), func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		version, err := strconv.Atoi(path[strings.LastIndex(path, string(os.PathSeparator))+1 : strings.LastIndex(path, ".")])
		if err != nil {
			return err
		}
		versions = append(versions, version)

		return nil
	})

	if err != nil {
		return err
	}

	//排序
	for i := 0; i < len(versions)-1; i++ {
		for j := 0; j < len(versions)-i-1; j++ {
			if versions[j] < versions[j+1] {
				temp := versions[j]
				versions[j] = versions[j+1]
				versions[j+1] = temp
			}
		}
	}

	for i := 0; i < len(versions); i++ {
		fmt.Println("release " + strconv.Itoa(versions[i]))

		var changeLog models.ChangeLog

		err = util.LoadJSON(".doup"+string(os.PathSeparator)+"changelogs"+string(os.PathSeparator)+strconv.Itoa(versions[i])+".json", &changeLog)
		if err != nil {
			return err
		}

		t := time.Unix(changeLog.TimeStamp, 0)

		fmt.Println("Time:", t)
		fmt.Println("Version:", changeLog.Version)
		fmt.Println("Changes:\n- " + strings.ReplaceAll(changeLog.Changes, "\n", "\n- "))
		fmt.Println()

		if (i-1)%2 == 0 && i != len(versions)-1 {
			line := ""
			fmt.Print("(type 'q' to quit or other to continue):")
			_, _ = fmt.Scanln(&line)
			if line == "q" {
				break
			}
		}
	}

	return nil
}
