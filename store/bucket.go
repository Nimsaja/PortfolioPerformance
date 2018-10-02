package store

import (
	"context"
	"fmt"

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
	wc.ContentType = "text/plain"

	if _, err := wc.Write([]byte("Yeah, I am writing to a bucket :-D")); err != nil {
		return fmt.Errorf("createFile: unable to write data to bucket %v, file %q: %v", bucket, fileName, err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("createFile: unable to close bucket %v, file %q: %v", bucket, fileName, err)
	}

	return nil
}
