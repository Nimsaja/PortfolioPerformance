package store

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// File store for save data to File
type File struct {
	path string
}

// HistoricalData time and quote value load from file
type HistoricalData struct {
	Time  int
	Value float32
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
func (file File) Load(quote float32) error {
	//read in file
	f, err := os.OpenFile(file.path, os.O_RDWR, 0600)
	if err != nil {
		return fmt.Errorf("Can not open file: %s, %v", file.path, err)
	}

	defer f.Close()

	getHistoricalData(f)

	return nil
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

		a = append(a, HistoricalData{t, float32(v)})
	}

	return a, nil
}