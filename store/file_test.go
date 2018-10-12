package store

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestCreatePath(t *testing.T) {
	s := " Test Me "

	p := createPath(s)

	if p != "testme.txt" {
		t.Errorf("Expected path: %v, got %v", "testme", p)
	}
}

func TestGetData(t *testing.T) {
	s := `[{"time":1538171940,"value":3328.5918,"diff":200},
	{"time":1538172010,"value":3114.1509,"diff":215}]`

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

func data() []Data {
	data := []Data{
		Data{Time: 1538431140, Value: 3798.3943, Diff: 191.1499}, //1.10.18
		Data{Time: 1538517540, Value: 3798.3943, Diff: 191.1499}, //2.10.18
		Data{Time: 1538603940, Value: 3797.939, Diff: 190.69458}, //3.10.18
		Data{Time: 1538690340, Value: 3704.2249, Diff: 96.98047}} //4.10.18

	for i, d := range data {
		d.TimeHuman = time.Unix(int64(d.Time), 0)
		data[i] = d
	}

	return data
}

func TestAppendToListNewData(t *testing.T) {
	newData := Data{Value: 4000, Diff: 100}
	regMarketTime := 1538776740 //5.10.18

	a := AppendToList(data(), newData, int64(regMarketTime))

	if len(a) != 5 {
		t.Errorf("New Data should be appended to list. Expected length of %v, got %v", 5, len(a))
	}

	dLast := a[4]
	if dLast.TimeHuman.Day() != 5 {
		t.Errorf("Last entry should be for day %v, got %v", 5, dLast.TimeHuman)
	}
	if dLast.Value != 4000 {
		t.Errorf("Last entry value should be %v, got %v", 4000, dLast.Value)
	}
	equal := reflect.DeepEqual(data(), a[:len(data())])
	if !equal {
		t.Errorf("Expected to be both arrays to be equal: \n%v \n!= \n%v", data(), a[:len(data())])
	}
}

func TestAppendToListDuplicatedData(t *testing.T) {
	newData := Data{Value: 4000, Diff: 100} //updated data for 4.10.18
	regMarketTime := 1538690340             //4.10.18

	a := AppendToList(data(), newData, int64(regMarketTime))

	if len(a) != 4 {
		t.Errorf("New Data should have overriden last entry of list. Expected length of %v, got %v", 4, len(a))
	}

	dLast := a[3]
	if dLast.TimeHuman.Day() != 4 {
		t.Errorf("Last entry should be for day %v, got %v", 4, dLast.TimeHuman)
	}
	if dLast.Value != 4000 {
		t.Errorf("Last entry value should be %v, got %v", 4000, dLast.Value)
	}
	l := len(data()) - 1
	equal := reflect.DeepEqual(data()[:l], a[:l])
	if !equal {
		t.Errorf("Expected to be both arrays to be equal: \n%v \n!= \n%v", data()[:l], a[:l])
	}
}
