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

//NewFile ...
func NewFile(s string) File {
	return File{
		path: createPath(s),
	}
}

//Save store quote into file
func (file File) Save(c context.Context, quote float32, buy float32) error {
	f, err := os.OpenFile(file.path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("Can not open file: %s, %v", file.path, err)
	}

	defer f.Close()

	//get data from file first
	data, err := getData(f)
	if err != nil {
		return err
	}

	t := calcStoreTime()
	data = append(data, Data{Time: int(t), TimeHuman: time.Unix(t, 0), Value: quote, Diff: quote - buy})
	s, err := convert2JSON(data)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, s)

	return err
}

//Load ...
func (file File) Load(c context.Context) ([]Data, error) {
	//read in file
	f, err := os.OpenFile(file.path, os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("Can not open file: %s, %v", file.path, err)
	}

	defer f.Close()

	return getData(f)
}

func getData(r io.Reader) (data []Data, err error) {
	// var d Data
	// prevTimes := make(map[int]struct{})

	byteValue, _ := ioutil.ReadAll(r)
	json.Unmarshal(byteValue, &data)

	// scanner := bufio.NewScanner(r)
	// for scanner.Scan() {
	// 	err := json.Unmarshal([]byte(scanner.Text()), &d)
	// 	if err != nil {
	// 		return data, fmt.Errorf("Can not unmarshal json. %v", err)
	// 	}

	// 	//check if this time already exists in map
	// 	_, exists := prevTimes[d.Time]
	// 	if exists {
	// 		continue
	// 	}
	// 	prevTimes[d.Time] = struct{}{}

	// 	data = append(data, d)
	// }
	return data, nil
}

//GetTime gets the time as time format
func getTime(t int) time.Time {
	return time.Unix(int64(t), 0)
}
