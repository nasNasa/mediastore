package storage

import "go.mongodb.org/mongo-driver/bson/primitive"

type File struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string
	Type       string        `bson:"type,omitempty"`
	Address    interface{}   `bson:"address,omitempty"`
	size       uint64        `bson:"size,omitempty"`
	UserFolder []interface{} `bson:"user_folder,omitempty"`
}
