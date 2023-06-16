package fileservice

import (
	"bytes"
	"context"
	"mediaStorer/param/storage"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	fileRepository FileRepository
	userRepository UserRepository
}

func New(fileRepository FileRepository, userRepository UserRepository) Service {
	return Service{fileRepository, userRepository}
}

type FileRepository interface {
	Insert(ctx context.Context, file storage.File) error
	UploadFile(file multipart.File, fileName string, userId uint) (primitive.ObjectID, error)
	Download(ctx context.Context, file_id primitive.ObjectID) (bytes.Buffer, error)
}

type UserRepository interface {
	AddObjectIdToUser(objectId primitive.ObjectID, userId uint) error
}
