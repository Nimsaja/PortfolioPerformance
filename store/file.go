package store

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

// File store for save data to File
type File struct {
	path string
}

// Data time and quote value load from file
type Data struct {
	Time      int       `json:"time"`
	TimeHuman time.Time `json:"timehuman"` //need to check if this can be done on the client
	Value     float32   `json:"value"`
	Diff      float32   `json:"diff"`
}

var today = time.Now().Day()
var skipDates = make(map[int]struct{})

//Save store quote into file
func (file File) Save(c context.Context, quote float32, buy float32, regMTime int64) error {
	f, err := os.OpenFile(file.path, os.O_RDONLY, 0600)
	if err != nil {
		return fmt.Errorf("Can not open file: %s, %v", file.path, err)
	}

	defer f.Close()

	//get data from file first
	data, err := getData(f)
	if err != nil {
		return err
	}

	newData := appendToList(data, Data{Value: quote, Diff: quote - buy}, regMTime)

	s, err := convert2JSON(newData)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file.path, []byte(s), 0644)
	if err != nil {
		return err
	}

	return err
}

//Load ...
func (file File) Load(c context.Context) ([]Data, error) {
	//read in file
	f, err := os.OpenFile(file.path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("Can not open file: %s, %v", file.path, err)
	}

	defer f.Close()

	return getData(f)
}

//append last data to list - override if we have already this date in the list
func appendToList(data []Data, d Data, regMTime int64) []Data {
	//skip this if regularMarketTime is not today -> weekends, public holidays
	if today > time.Unix(regMTime, 0).Day() {
		skipDates[today] = struct{}{}
		return data
	}

	//closure date
	t := calcStoreTime(regMTime)

	//try yesterday, if not valid, go back one day etc.
	skip := true
	for skip {
		_, skip = skipDates[t.Day()]

		if skip {
			t = t.Add(time.Duration(-1) * time.Hour * 24)
			continue
		}

		d.Time = int(t.Unix())
		d.TimeHuman = t

		//clear map
		skipDates = make(map[int]struct{})
	}

	a := make([]Data, len(data))
	copy(a, data)

	//check if this is already in list, can only be the last element -> override the values
	if a[len(a)-1].TimeHuman.Day() == t.Day() {
		a[len(a)-1] = d
	} else {
		a = append(a, d)
	}

	return a
}

func getData(r io.Reader) ([]Data, error) {
	byteValue, _ := ioutil.ReadAll(r)
	var res []Data
	json.Unmarshal(byteValue, &res)

	return res, nil
}
