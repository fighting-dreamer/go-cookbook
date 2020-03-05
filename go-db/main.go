package main

import (
	"fmt"
	"nipun.io/go-cookbook/go-db/db"
)

func main() {
	db.Conn = db.WillConnect()
	err := db.Conn.Ping()
	if err != nil {
		fmt.Println("nwkfnkfn")
	}
	db.CheckClose(db.Conn)
}