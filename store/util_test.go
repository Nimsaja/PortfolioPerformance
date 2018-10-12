package store

import (
	"encoding/json"
	"testing"
)

func TestConvert2JSON(t *testing.T) {
	o := "Test Text"

	r, err := convert2JSON(o)
	if err != nil {
		t.Errorf("No error expected! %v", err)
	}

	var s string
	err = json.Unmarshal([]byte(r), &s)

	if s != o {
		t.Errorf("Expected %v, got %v", "Test Text", r)
	}
}
