package store

import (
	"encoding/json"
	"fmt"
	"strings"
)

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
