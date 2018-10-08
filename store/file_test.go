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

func TestStoreValuesInSortedArray(t *testing.T) {
	s := []Data{{Time: 1538431140, Value: 3798.3943, Diff: 100.001},
		{Time: 1538517540, Value: 3798.3943, Diff: 100.002},
		{Time: 1538863140, Value: 3704.2249, Diff: 100.006},
		{Time: 1538603940, Value: 3797.939, Diff: 100.003},
		{Time: 1538690340, Value: 3704.2249, Diff: 100.004},
		{Time: 1538776740, Value: 3704.2249, Diff: 100.005},
		{Time: 1538949540, Value: 3704.2249, Diff: 100.007}}

	times := make(map[int]Data, 7)
	for _, d := range s {
		times[d.Time] = d
	}

	data := storeValuesInSortedArray(times)

	var c float32
	for i := 0; i < 7; i++ {
		c = float32(i+1)/1000.0 + 100

		if data[i].Diff != c {
			t.Errorf("Expected diff of %v, got %v ", c, data[i].Diff)
		}
	}
}
