package store

import (
	"encoding/json"
	"fmt"
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

func convert2JSON(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return string(""), fmt.Errorf("Can not convert to JSON: %v", err)
	}

	return string(b), nil
}

func jsonData(quote, buy float32) (string, error) {
	t := calcStoreTime()
	data := Data{Time: int(t), TimeHuman: time.Unix(t, 0), Value: quote, Diff: quote - buy}
	s, err := convert2JSON(data)
	if err != nil {
		return "", err
	}
	return s, nil
}
