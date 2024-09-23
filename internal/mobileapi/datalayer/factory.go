package datalayer

import "pillsreminder/internal/mobileapi/datalayer/postgres"

type StorageFactory struct{}

func NewStorageFactory() *StorageFactory {
	return &StorageFactory{}
}

func (f *StorageFactory) GetUserRepo() IUserRepo {
	return postgres.NewUserRepo()
}
