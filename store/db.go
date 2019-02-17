package store

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
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
	c = context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := "portfolio-218213"

	// Creates a client.
	client, err := datastore.NewClient(c, projectID)
	if err != nil {
		return fmt.Errorf("Failed to create client: %v", err)
	}

	// creates checkTime, to filter out last saved data
	l, err := time.LoadLocation("Local")
	if err != nil {
		return fmt.Errorf("can't find location %v ", err)
	}
	checkTime := time.Unix(regMTime, 0)
	checkTime = time.Date(checkTime.Year(), checkTime.Month(), checkTime.Day(), 0, 0, 0, 0, l)

	// get entries with a time later then checkTime
	query := datastore.NewQuery("SaveData").Filter("time >=", checkTime)
	noData := true

	it := client.Run(c, query)
	for {
		var d SaveData
		k, err := it.Next(&d)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Error fetching next data: %v", err)
		}
		noData = false

		//overwrite with latest value
		d.TodaySum = quote
		d.Diff = quote - buy
		d.Time = time.Unix(regMTime, 0)
		if _, err := client.Put(c, k, &d); err != nil {
			return fmt.Errorf("Failed to save quote: %v", err)
		}
	}

	// first data of the day! :-D
	if noData {
		// Sets the kind for the new entity.
		kind := "SaveData"

		// Creates a SaveData instance.
		data := &SaveData{
			TodaySum: quote,
			Diff:     quote - buy,
			Time:     time.Unix(regMTime, 0),
		}

		key := datastore.IncompleteKey(kind, nil)

		if _, err := client.Put(c, key, data); err != nil {
			return fmt.Errorf("Failed to save quote: %v", err)
		}
	}

	return nil
}

//Load loads values from database
func (f Database) Load(c context.Context) ([]Data, error) {
	return nil, nil
}
