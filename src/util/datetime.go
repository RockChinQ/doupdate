package util

import "time"

func GetNowDateTime() int {
	t := time.Now()
	return t.Year()*10000 + int(t.Month())*100 + t.Day()
}
