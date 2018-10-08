package store

import (
	"strings"
	"testing"
)

func TestCreatePath(t *testing.T) {
	s := " Test Me "

	p := createPath(s)

	if p != "testme.txt" {
		t.Errorf("Expected path: %v, got %v", "testme", p)
	}
}

func checkData(s string, t *testing.T) {
	a, err := getData(strings.NewReader(s))
	if err != nil {
		t.Errorf("No Error expected %v", err)
	}
	if len(a) != 2 {
		t.Errorf("Expected an array length of %v, got %v", 2, len(a))
	}
	if a[0].Time != 1538171940 {
		t.Errorf("Expected time %v, got %v", 1538171940, a[0].Time)
	}
	if a[0].Value != 3328.5918 {
		t.Errorf("Expected value %v, got %v", 3328.5918, a[0].Value)
	}
	if a[0].Diff != 200 {
		t.Errorf("Expected doff %v, got %v", 200, a[0].Time)
	}
	if a[1].Time != 1538172010 {
		t.Errorf("Expected time %v, got %v", 1538172010, a[1].Time)
	}
	if a[1].Value != 3114.1509 {
		t.Errorf("Expected value %v, got %v", 3114.1509, a[1].Value)
	}
	if a[1].Diff != 215 {
		t.Errorf("Expected value %v, got %v", 215, a[1].Diff)
	}
}

func TestGetData(t *testing.T) {
	s := `[{"time":1538171940,"value":3328.5918,"diff":200},
	{"time":1538172010,"value":3114.1509,"diff":215}]`

	checkData(s, t)
}

func TestRemovePreviousDuplicatesFromData(t *testing.T) {
	s := `[{"time":1538171940,"value":3114.1509,"diff":215},
	{"time":1538172010,"value":3114.1509,"diff":215},
	{"time":1538171940,"value":3328.5918,"diff":200}]`

	checkData(s, t)
}

func TestSortedDataArray(t *testing.T) {
	s := `[{"time":1538690340,"timehuman":"2018-10-04T21:59:00Z","value":3704.2249,"diff":100.004},
	{"time":1538949540,"timehuman":"2018-10-07T21:59:00Z","value":3593.6108,"diff":100.007},
	{"time":1538603940,"timehuman":"2018-10-03T21:59:00Z","value":3797.939,"diff":100.003},
	{"time":1538863140,"timehuman":"2018-10-06T21:59:00Z","value":3704.2249,"diff":100.006},
	{"time":1538776740,"timehuman":"2018-10-05T21:59:00Z","value":3704.2249,"diff":100.005},
	{"time":1538517540,"timehuman":"2018-10-02T21:59:00Z","value":3798.3943,"diff":100.002},
	{"time":1538431140,"timehuman":"2018-10-01T21:59:00Z","value":3798.3943,"diff":100.001}]`

	a, err := getData(strings.NewReader(s))
	if err != nil {
		t.Errorf("No Error expected %v", err)
	}

	var c float32
	for i := 0; i < 7; i++ {
		c = float32(i+1)/1000.0 + 100

		if a[i].Diff != c {
			t.Errorf("Expected diff of %v, got %v ", c, a[i].Diff)
		}
	}
}
