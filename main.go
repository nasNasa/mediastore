package main

import (
	_ "github.com/go-sql-driver/mysql"
	"mediaStorer/delivery/httpServer"
	"mediaStorer/repository/mysql"
	"mediaStorer/repository/mysql/mysqluser"
	"mediaStorer/service/userservice"
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

	server1 := httpServer.New(usersvc)
	server1.Serve()
}
