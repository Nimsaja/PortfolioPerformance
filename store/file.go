package store

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// File store for save data to File
type File struct {
	path string
}

// HistoricalData time and quote value load from file
type HistoricalData struct {
	Time      int       `json:"time"`
	TimeHuman time.Time `json:"timehuman"` //need to check if this can be done on the client
	Value     float32   `json:"value"`
}

//NewFile ...
func NewFile(s string) File {
	return File{
		path: createPath(s),
	}
}

func createPath(s string) string {
	path := strings.Replace(s, " ", "", -1)
	path = strings.ToLower(path)
	return path + ".txt"
}

//Save store quote into file
func (file File) Save(quote float32) error {
	//append to output file
	f, err := os.OpenFile(file.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("Can not open file: %s, %v", file.path, err)
	}

	defer f.Close()

	s := fmt.Sprintf("%v, %v", calcStoreTime(), quote)
	_, err = fmt.Fprintln(f, s)

	return err
}

//Load ...
func (file File) Load() ([]HistoricalData, error) {
	//read in file
	f, err := os.OpenFile(file.path, os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("Can not open file: %s, %v", file.path, err)
	}

	defer f.Close()

	return getHistoricalData(f)
}

func getHistoricalData(r io.Reader) ([]HistoricalData, error) {
	a := make([]HistoricalData, 0)
	var s []string
	var v float64
	prevTimes := make(map[int]struct{})

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		b := scanner.Text()
		s = strings.Split(b, ", ")

		//get time
		t, err := strconv.Atoi(s[0])
		if err != nil {
			return nil, fmt.Errorf("Error parsing time %v", err)
		}

		//check if this time already exists in map
		_, exists := prevTimes[t]
		if exists {
			continue
		}
		prevTimes[t] = struct{}{}

		//get value
		v, err = strconv.ParseFloat(s[1], 32)
		if err != nil {
			return nil, fmt.Errorf("Error parsing quote value %v", err)
		}

		a = append(a, HistoricalData{Time: t, TimeHuman: getTime(t), Value: float32(v)})
	}

	return a, nil
}

//GetTime gets the time as time format
func getTime(t int) time.Time {
	return time.Unix(int64(t), 0)
}
