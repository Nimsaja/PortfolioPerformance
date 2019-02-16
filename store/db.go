package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

// SaveData structure for datastore
type SaveData struct {
	TodaySum float32   `datastore:"profit"`
	Diff     float32   `datastore:"sum"`
	Time     time.Time `datastore:"time"`
}

// NewDB creates a new datastore
func NewDB() (ctx context.Context, client *datastore.Client) {
	ctx = context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := "portfolio-218213"

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return ctx, client
}

// SaveQuote saves a quote
func SaveQuote(ctx context.Context, client *datastore.Client, quote float32, diff float32, t time.Time) {
	// Sets the kind for the new entity.
	kind := "SaveData"

	// Creates a SaveData instance.
	data := &SaveData{
		TodaySum: quote,
		Diff:     diff,
		Time:     t,
	}

	key := datastore.IncompleteKey(kind, nil)

	if _, err := client.Put(ctx, key, data); err != nil {
		log.Fatalf("Failed to save quote: %v", err)
	}

	fmt.Printf("Saved %v: %v\n", key, data)

	// gets Entry with date from last saved date (today - 1)
	l, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Printf("can't find location %v ", err)
	}
	checkTime := time.Now()
	checkTime = time.Date(checkTime.Year(), checkTime.Month(), checkTime.Day()-1,
		0, 0, 0, 0, l)

	fmt.Printf("checkTime %v\n", checkTime)

	// get entries with a time later then checkTime
	query := datastore.NewQuery("SaveData").Filter("time >=", checkTime)

	it := client.Run(ctx, query)
	for {
		var d SaveData
		_, err := it.Next(&d)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next data: %v", err)
		}
		fmt.Printf("SavedData for Date %v : %v\n", checkTime, d)
	}
}
