package store

import (
	"context"
)

//StorageService interface
type StorageService interface {
	Save(c context.Context, quote float32, buy float32) error
	Load(c context.Context) ([]Data, error)
}

//New factory to create an implementation of a StorageService interface
func New(inCloud bool, name string) StorageService {
	p := createPath(name)
	var s StorageService
	if inCloud {
		s = Bucket{path: p}
	} else {
		s = File{path: p}
	}
	return s
}
