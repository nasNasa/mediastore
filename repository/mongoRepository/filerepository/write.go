package filerepository

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mediaStorer/param/storage"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *DB) Insert(ctx context.Context, file storage.File) error {
	res, err := d.conn.Collection.InsertOne(ctx, file)
	fmt.Println(res.InsertedID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (d *DB) UploadFile(file multipart.File, fileName string, userId uint) (primitive.ObjectID, error) {
	fmt.Println("file test")

	opts := options.GridFSBucket().SetName("fileRepository")
	bucket, err := gridfs.NewBucket(d.conn.Database(), opts)
	if err != nil {
		fmt.Println("createbucket error", err)
	}

	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{"metadata tag", "first"}})
	objectID, err := bucket.UploadFromStream(fileName, io.Reader(file), uploadOpts)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return objectID, nil
}

func (d *DB) Delete() {

}

func (d *DB) Download(ctx context.Context, id primitive.ObjectID) (bytes.Buffer, error) {

	opts := options.GridFSBucket().SetName("fileRepository")
	bucket, err := gridfs.NewBucket(d.conn.Database(), opts)
	if err != nil {
		fmt.Println("createbucket error", err)
	}

	fileBuffer := bytes.NewBuffer(nil)
	if _, err := bucket.DownloadToStream(id, fileBuffer); err != nil {
		return *bytes.NewBuffer(nil), err
	}

	return *fileBuffer, nil
}
