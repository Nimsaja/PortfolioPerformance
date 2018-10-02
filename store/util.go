package store

import (
	"log"
	"strings"
	"time"
)

func calcStoreTime() int64 {
	//date should be the close time from yesterday - say 23:59
	d := time.Now().Add(time.Duration(-1) * time.Hour * 24)
	l, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		log.Printf("Can not load location, %v", err)
	}
	d = time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 0, 0, l)

	return d.Unix()
}

func createPath(s string) string {
	path := strings.Replace(s, " ", "", -1)
	path = strings.ToLower(path)
	return path + ".txt"
}
