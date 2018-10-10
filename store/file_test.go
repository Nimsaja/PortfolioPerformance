package store

import (
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
	today = 6
	newData := Data{Value: 3704.2249, Diff: 96.98047} //yesterdays data - 5.10.18
	regMarketTime := 1538863140                       //6.10.18

	a := appendToList(data(), newData, int64(regMarketTime))

	if len(a) != 5 {
		t.Errorf("New Data should be appended to list. Expected length of %v, got %v", 5, len(a))
	}

	dLast := a[4]
	if dLast.TimeHuman.Day() != 5 {
		t.Errorf("Last entry should be for day %v, got %v", 5, dLast.TimeHuman)
	}
	if dLast.Value != 3704.2249 {
		t.Errorf("Last entry value should be %v, got %v", 3704.2249, dLast.Value)
	}

	if len(skipDates) != 0 {
		t.Errorf("Expected a length of 0, got %v ", len(skipDates))
	}
}
func TestAppendToListDuplicatedData(t *testing.T) {
	today = 5
	newData := Data{Value: 4000, Diff: 100} //yesterdays updated data - 4.10.18
	regMarketTime := 1538776740             //5.10.18

	a := appendToList(data(), newData, int64(regMarketTime))

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
	if len(skipDates) != 0 {
		t.Errorf("Expected a length of 0, got %v ", len(skipDates))
	}
}
func TestAppendToListSkipWeekends(t *testing.T) {
	today = 6
	newData := Data{Value: 4000, Diff: 100} //yesterdays updated data - 4.10.18
	regMarketTime := 1538776740             //5.10.18 - not today!!!!!

	a := appendToList(data(), newData, int64(regMarketTime))

	if len(a) != 4 {
		t.Errorf("New Data from weekend should be skipped. Expected length of %v, got %v", 4, len(a))
	}

	dLast := a[3]
	if dLast.TimeHuman.Day() != 4 {
		t.Errorf("Last entry should be for day %v, got %v", 4, dLast.TimeHuman)
	}
	if dLast.Value != 3704.2249 {
		t.Errorf("Last entry value should not be overriden! Expected %v, got %v", 3704.2249, dLast.Value)
	}
	_, exists := skipDates[6]
	if !exists {
		t.Errorf("Should store day %v to skipDates. Got %v ", 6, skipDates)
	}
}

func TestAppendToListAfterWeekend(t *testing.T) {
	//Friday 5.10.
	today = 5
	newData := Data{Value: 5000, Diff: 100} //yesterdays updated data - 4.10.18
	regMarketTime := 1538776740             //5.10.18

	a := appendToList(data(), newData, int64(regMarketTime))

	if len(a) != 4 {
		t.Errorf("Expected length of %v, got %v", 4, len(a))
	}

	dLast := a[3]
	if dLast.TimeHuman.Day() != 4 {
		t.Errorf("Last entry should be for day %v, got %v", 4, dLast.TimeHuman)
	}
	if dLast.Value != 5000 {
		t.Errorf("Last entry value should be overriden! Expected %v, got %v", 5000, dLast.Value)
	}
	if len(skipDates) != 0 {
		t.Errorf("There should be no skip dates stored. Got %v ", skipDates)
	}

	//Saturday 6.10.
	today = 6
	newData = Data{Value: 6000, Diff: 100} //yesterdays updated data - still for the 4.10.18
	regMarketTime = 1538776740             //still 5.10.18 as no new data are available

	a = appendToList(a, newData, int64(regMarketTime))

	if len(a) != 4 {
		t.Errorf("Saturday values should be skipped! Expected length of %v, got %v", 4, len(a))
	}

	dLast = a[3]
	if dLast.TimeHuman.Day() != 4 {
		t.Errorf("Last entry should be for day %v, got %v", 4, dLast.TimeHuman)
	}
	if dLast.Value != 5000 {
		t.Errorf("Last entry value should still be from the day before! Expected %v, got %v", 5000, dLast.Value)
	}

	_, exists := skipDates[6]
	if !exists {
		t.Errorf("Should store Saturday %v to skipDates. Got %v ", 6, skipDates)
	}

	//Sunday 7.10.
	today = 7
	newData = Data{Value: 7000, Diff: 100} //yesterdays updated data - still for the 4.10.18
	regMarketTime = 1538776740             //still 5.10.18 as no new data are available

	a = appendToList(a, newData, int64(regMarketTime))

	if len(a) != 4 {
		t.Errorf("Sunday values should be skipped! Expected length of %v, got %v", 4, len(a))
	}

	dLast = a[3]
	if dLast.TimeHuman.Day() != 4 {
		t.Errorf("Last entry should be for day %v, got %v", 4, dLast.TimeHuman)
	}
	if dLast.Value != 5000 {
		t.Errorf("Last entry value should still be from the day before! Expected %v, got %v", 5000, dLast.Value)
	}

	_, exists = skipDates[7]
	if !exists {
		t.Errorf("Should store Sunday %v to skipDates. Got %v ", 7, skipDates)
	}
	_, exists = skipDates[6]
	if !exists {
		t.Errorf("Should still have Saturday %v in skipDates. Got %v ", 6, skipDates)
	}

	//Monday 8.10.
	today = 8
	newData = Data{Value: 8000, Diff: 100} //yesterdays updated data - for 5.10.18
	regMarketTime = 1539035940             //8.10.18

	a = appendToList(a, newData, int64(regMarketTime))

	if len(a) != 5 {
		t.Errorf("Expected length of %v, got %v", 5, len(a))
	}

	dLast = a[4]
	if dLast.TimeHuman.Day() != 5 {
		t.Errorf("Last entry should be for day %v, got %v", 5, dLast.TimeHuman)
	}
	if dLast.Value != 8000 {
		t.Errorf("Expected %v, got %v", 8000, dLast.Value)
	}

	if len(skipDates) != 0 {
		t.Errorf("There should be no skip dates stored. Got %v ", skipDates)
	}

	//Tuesday 9.10.
	today = 9
	newData = Data{Value: 9000, Diff: 100} //yesterdays updated data - for 8.10.18
	regMarketTime = 1539122340             //9.10.18

	a = appendToList(a, newData, int64(regMarketTime))

	if len(a) != 6 {
		t.Errorf("Expected length of %v, got %v", 5, len(a))
	}

	dLast = a[5]
	if dLast.TimeHuman.Day() != 8 {
		t.Errorf("Last entry should be for day %v, got %v", 8, dLast.TimeHuman)
	}
	if dLast.Value != 9000 {
		t.Errorf("Expected %v, got %v", 9000, dLast.Value)
	}

	if len(skipDates) != 0 {
		t.Errorf("There should be no skip dates stored. Got %v ", skipDates)
	}
}
