package fileservice

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Service) UploadFile(file multipart.File, fileName string, userId uint) (primitive.ObjectID, error) {
	objectId, err := s.fileRepository.UploadFile(file, fileName, userId)
	if err != nil {
		return primitive.NewObjectID(), err
	}
	fmt.Println("objectid", objectId)

	// Todo add userid
	s.userRepository.AddObjectIdToUser(objectId, 123)

	return objectId, nil
}

func (s Service) DownloadFile(file_id primitive.ObjectID, userId uint) (bytes.Buffer, error) {
	file, err := s.fileRepository.Download(context.TODO(), file_id)
	if err != nil {
		return *bytes.NewBuffer(nil), fmt.Errorf("downloading error", err)
	}

	// Todo add userid
	// s.userRepository.AddObjectIdToUser(objectId, 123)

	return file, nil
}

//
//func (s Service) DeleteFile(filerepository storage.File) error {
//	//check access
//	return nil
//}
