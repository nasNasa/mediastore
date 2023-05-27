package main

import (
	_ "github.com/go-sql-driver/mysql"
	"mediaStorer/config"
	"mediaStorer/delivery/httpServer"
	"mediaStorer/repository/mysql"
	"mediaStorer/repository/mysql/mysqluser"
	"mediaStorer/service/authentication"
	"mediaStorer/service/userservice"
	"time"
)

func main() {

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
