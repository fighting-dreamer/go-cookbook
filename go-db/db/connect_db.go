package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var Conn *sqlx.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypass"
	dbname   = "mydb"
	dataSourceName = "postgres://myuser:mypass@localhost:5432/mydb"
)

func WillNotConnect() *sqlx.DB {
	fmt.Println("in WillNotConnect function : ")
	driver := "postgres"
	psqlInfo := dataSourceName
	db, err := sqlx.Open(driver, psqlInfo)
	if err != nil {
		log.Fatal("there seem to be an err while connecting to DB", err.Error())
	}
	return db
}

func WillConnect() *sqlx.DB{
	fmt.Println("in connect function : ")
	driver := "postgres"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open(driver, psqlInfo)
	if err != nil {
		log.Fatal("there seem to be an err while connecting to DB", err.Error())
	}
	return db
}

func CheckClose(conn *sqlx.DB) {
	err := conn.Close()
	if err != nil {
		log.Fatal("there was an error while trying to close ", err.Error())
	}
}

