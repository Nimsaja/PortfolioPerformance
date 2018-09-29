package store

import (
	"fmt"
	"os"
	"strings"
)

// File store for save data to File
type File struct {
	path string
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
	return nil
}
