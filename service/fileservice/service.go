package fileservice

import "mediaStorer/param/storage"

type Service struct {
	repository Repository
}

func New(repository Repository) Service {
	return Service{repository: repository}
}

type Repository interface {
	WriteFileToDatabase(file storage.File) error
}
