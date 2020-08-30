package utils

import (
	"fmt"
	"time"
)

func FormatDate(dateTime string) string {
	layout := "2006-01-02T15:04"
	t, err := time.Parse(layout, dateTime)

	if err != nil {
		fmt.Println(err)
	}

	runes := []rune(fmt.Sprint(t))

	return string(runes[0:16])
}

func GetGDSharedUrl(fileId string) (urlfile string) {
	return fmt.Sprintf("https://drive.google.com/file/d/%s/view?usp=sharing", fileId)
}