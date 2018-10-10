package store

import (
	"encoding/json"
	"testing"
)

func TestCalcStoreTime(t *testing.T) {
	st := calcStoreTime(1538603940)

	if st.Month() != 10 {
		t.Errorf("Expected month %v, got %v ", 10, st.Month())
	}
	if st.Day() != 2 {
		t.Errorf("Expected day %v, got %v ", 2, st.Day())
	}
	if st.Hour() != 23 {
		t.Errorf("Expected hour of %v, got %v", 23, st.Hour())
	}
}

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
