package store

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/appengine/datastore"
)

const (
	// KIND of datastore
	KIND = "SaveData"
)

//Database stores the name of the owner
type Database struct {
	owner string
}

// SaveData structure for datastore
type SaveData struct {
	TodaySum float32   `datastore:"profit"`
	Diff     float32   `datastore:"sum"`
	Time     time.Time `datastore:"time"`
}

//Save saves values to database
func (f Database) Save(c context.Context, quote float32, buy float32, regMTime int64) error {
	// Creates a SaveData instance.
	data := &SaveData{
		TodaySum: quote,
		Diff:     quote - buy,
		Time:     time.Unix(regMTime, 0),
	}

	// creates checkTime, to filter out last saved data
	l, err := time.LoadLocation("Local")
	if err != nil {
		return fmt.Errorf("can't find location %v ", err)
	}
	checkTime := time.Unix(regMTime, 0)
	checkTime = time.Date(checkTime.Year(), checkTime.Month(), checkTime.Day(), 0, 0, 0, 0, l)

	// get entries with a time later then checkTime
	q := datastore.NewQuery(KIND).Filter("time >=", checkTime).KeysOnly()

	keys, err := q.GetAll(c, nil)
	if err != nil {
		return fmt.Errorf("Err by datastore.GetAll: %v", err)
	}

	s := len(keys)

	if s > 1 {
		return fmt.Errorf("To many results in database %v", l)
	}

	var k *datastore.Key

	// first data of the day! :-D
	if len(keys) == 0 {
		k = datastore.NewIncompleteKey(c, KIND, nil)
	} else {
		// overwrite old value with new one
		k = keys[0]
	}

	if _, err := datastore.Put(c, k, data); err != nil {
		return fmt.Errorf("Failed to save quote: %v", err)
	}

	return nil
}

//Load loads values from database
func (f Database) Load(c context.Context) ([]Data, error) {
	return nil, nil
}
