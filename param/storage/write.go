package storage

import "mime/multipart"

type File struct {
	//ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	Type string `bson:"type,omitempty"`
	//Address    interface{}        `bson:"address,omitempty"`
	Size uint64         `bson:"size,omitempty"`
	File multipart.File `bson:"file,omitempty"`
	//UserFolder interface{}        `bson:"user_folder,omitempty"`
}
