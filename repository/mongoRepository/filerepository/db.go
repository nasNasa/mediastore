package filerepository

import (
	"mediaStorer/repository/mongoRepository"
)

type DB struct {
	conn *mongoRepository.Db
}

func New(conn *mongoRepository.Db) *DB {
	return &DB{
		conn: conn,
	}
}
