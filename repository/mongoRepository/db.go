package mongoRepository

import "go.mongodb.org/mongo-driver/mongo"

type Db struct {
	config Config
	*mongo.Collection
}

func New(config Config, collection mongo.Collection) Db {
	return Db{config, &collection}
}

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}
