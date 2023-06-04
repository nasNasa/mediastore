package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mediaStorer/config"
	"mediaStorer/delivery/httpServer"
	"mediaStorer/repository/mongoRepository"
	"mediaStorer/repository/mysql"
	"mediaStorer/repository/mysql/mysqluser"
	"mediaStorer/service/authentication"
	"mediaStorer/service/userservice"
	"time"
)

func main() {

	mongoCfg := mongoRepository.Config{
		Username: "hasan",
		Password: "12345678",
		Host:     "localhost",
		Port:     "27017",
		DBName:   "storage",
	}

	//co := options.Client().ApplyURI("mongodb://hasan:12345678@localhost:27017")
	////
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://hasan:12345678@localhost:27017"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("ass", err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Initialize a new instance of application containing the dependencies.
	mongoRepository.New(&mongodb.BookingModel{
		C: client.Database(*mongoDatabase).Collection("bookings"),
	})

	cnfg := mysql.Config{
		Username: "mediastore",
		Password: "mediastoret0lk2o20",
		Host:     "localhost",
		Port:     3308,
		DBName:   "mediastore_db",
	}

	mysqlserver := mysql.New(cnfg)
	mysqlRepository := mysqluser.New(mysqlserver)
	usersvc := userservice.New(mysqlRepository)

	authConfig := authentication.Config{
		SignKey:               config.JwtSignKey,
		AccessExpirationTime:  7 * time.Hour,
		RefreshExpirationTime: 27 * time.Hour,
		AccessSubject:         "access",
		RefreshSubject:        "refresh",
	}
	authSvc := authentication.New(authConfig)

	server1 := httpServer.New(usersvc, authSvc)
	server1.Serve()
}
