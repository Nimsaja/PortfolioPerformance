package store

import (
	"testing"
	"time"
)

func TestCalcStoreTime(t *testing.T) {
	st := time.Unix(calcStoreTime(), 0)
	ct := time.Now().Add(time.Duration(-1) * time.Hour * 24)

	if st.Month() != ct.Month() {
		t.Errorf("Expected month %v, got %v ", ct.Month(), st.Month())
	}
	if st.Day() != ct.Day() {
		t.Errorf("Expected day %v, got %v ", ct.Day(), st.Day())
	}
	if st.Hour() != 23 {
		t.Errorf("Expected hour of %v, got %v", 23, st.Hour())
	}
}
