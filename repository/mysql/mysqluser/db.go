package mysqluser

import "mediaStorer/repository/mysql"

type DB struct {
	conn *mysql.DB
}

func New(conn *mysql.DB) *DB {
	return &DB{
		conn: conn,
	}
}
