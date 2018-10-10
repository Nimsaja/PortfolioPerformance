package store

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

func calcStoreTime(regMTime int64) time.Time {
	//date should be the close time from the day before - say 23:59
	t := time.Unix(regMTime, 0)
	d := t.Add(time.Duration(-1) * time.Hour * 24)
	l, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		log.Printf("Can not load location, %v", err)
	}
	d = time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 0, 0, l)

	return d
}

func createPath(s string) string {
	path := strings.Replace(s, " ", "", -1)
	path = strings.ToLower(path)
	return path + ".txt"
}

func convert2JSON(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return string(""), fmt.Errorf("Can not convert to JSON: %v", err)
	}

	return string(b), nil
}
