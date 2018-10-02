package store

import "context"

//StorageService interface
type StorageService interface {
	Save(c context.Context, quote float32, buy float32) error
	Load(c context.Context) ([]Data, error)
}
