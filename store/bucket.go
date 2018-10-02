package store

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/appengine/file"
)

//Bucket stores the path to the bucket
type Bucket struct {
	path string
}

//NewBucket new bucket with path
func NewBucket(s string) Bucket {
	return Bucket{
		path: createPath(s),
	}
}

//Save saves values to bucket
func (f Bucket) Save(c context.Context, quote float32, buy float32) error {
	//determine default bucket name
	bucketName, err := file.DefaultBucketName(c)
	if err != nil {
		return fmt.Errorf("failed to get default GCS bucket name: %v", err)
	}

	client, err := storage.NewClient(c)
	if err != nil {
		return fmt.Errorf("failed to get default GCS bucket name: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)

	fileName := f.path
	wc := bucket.Object(fileName).NewWriter(c)
	wc.ContentType = "application/json"

	s, err := jsonData(quote, buy)
	if err != nil {
		return err
	}

	if _, err := wc.Write([]byte(s)); err != nil {
		return fmt.Errorf("createFile: unable to write data to bucket %v, file %q: %v", bucket, fileName, err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("createFile: unable to close bucket %v, file %q: %v", bucket, fileName, err)
	}

	return nil
}

//Load loads values from bucket
func (f Bucket) Load(c context.Context) ([]Data, error) {
	// determine default bucket name
	bucketName, err := file.DefaultBucketName(c)
	if err != nil {
		log.Fatalf("failed to get default GCS bucket name: %v", err)
		return nil, err
	}

	client, err := storage.NewClient(c)
	if err != nil {
		log.Fatalf("failed to get default GCS bucket name: %v", err)
		return nil, err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)

	rc, err := bucket.Object(f.path).NewReader(c)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	return getData(rc)
}
