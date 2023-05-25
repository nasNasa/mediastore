package mysql

import (
	"database/sql"
	"fmt"
	"time"
)

type DB struct {
	config Config
	db     *sql.DB
}

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}

func (d *DB) Conn() *sql.DB {
	return d.db
}

func New(config Config) *DB {
	// parseTime=true changes the output type of DATE and DATETIME values to time.Time
	// instead of []byte / string
	// The date or datetime like 0000-00-00 00:00:00 is converted into zero value of time.Time
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		config.Username, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &DB{
		config: config,
		db:     db,
	}
}
