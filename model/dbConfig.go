package model

import (
	"database/sql"
	"udemyTraining/http/BasicServer/util"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

func DBInit() {
	log.Println("Opening DB connection...")
	DB, err = sql.Open("mysql", "administrator:administrator@tcp(testdb.c11gtkstiblb.ap-southeast-2.rds.amazonaws.com:3306)/test?charset=utf8")
	util.Check(err)
	defer DB.Close()

	err = DB.Ping()
	util.Check(err)
}

func DBClose()  {
	DB.Close()
}


