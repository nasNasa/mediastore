package main

import (
	"context"
	"fmt"
	"mediaStorer/config"
	"mediaStorer/delivery/httpServer"
	"mediaStorer/repository/mongoRepository"
	"mediaStorer/repository/mongoRepository/filerepository"
	"mediaStorer/repository/mongoRepository/userrepository"
	"mediaStorer/repository/mysql"
	"mediaStorer/repository/mysql/mysqluser"
	"mediaStorer/service/authentication"
	"mediaStorer/service/fileservice"
	"mediaStorer/service/userservice"
	"time"

	_ "github.com/go-sql-driver/mysql"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	mongoCfg := mongoRepository.Config{
		Username: "mdstore",
		Password: "12345678",
		Host:     "localhost",
		Port:     "27017",
		DBName:   "storage",
	}

	//co := options.Client().ApplyURI("mongodb://hasan:12345678@localhost:27017")
	////
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mdstore:12345678@localhost:27017"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("client", err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	mongoDbRepo := client.Database("mediastore_db")
	// Initialize a new instance of application containing the dependencies.
	fileMongoDb := mongoRepository.New(
		mongoCfg,
		*mongoDbRepo.Collection("fileRepository"),
	)
	userMongoDb := mongoRepository.New(
		mongoCfg,
		*mongoDbRepo.Collection("userRepository"),
	)

	fileMongoRepository := filerepository.New(&fileMongoDb)
	userMongoRepository := userrepository.New(&userMongoDb)

	fileSvc := fileservice.New(fileMongoRepository, userMongoRepository)

	cnfg := mysql.Config{
		Username: "mediastore",
		Password: "mediastoret0lk2o20",
		Host:     "localhost",
		Port:     3308,
		DBName:   "mediastore_db",
	}

	mysqlserver := mysql.New(cnfg)
	mysqlRepository := mysqluser.New(mysqlserver)

	authConfig := authentication.Config{
		SignKey:               config.JwtSignKey,
		AccessExpirationTime:  7 * time.Hour,
		RefreshExpirationTime: 27 * time.Hour,
		AccessSubject:         "access",
		RefreshSubject:        "refresh",
	}
	authSvc := authentication.New(authConfig)

	usersvc := userservice.New(mysqlRepository, authSvc)

	server1 := httpServer.New(authConfig, usersvc, authSvc, fileSvc)
	server1.Serve()
}
