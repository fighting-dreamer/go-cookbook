package main

import (
	"fmt"
	"github.com/lib/pq"
	"nipun.io/go-cookbook/go-db/db"
)

func main() {
	db.Conn = db.WillConnect()
	defer db.CheckClose(db.Conn)
	err := db.Conn.Ping()
	if err != nil {
		fmt.Println("nwkfnkfn")
	}
	dl := db.DAO_Layer{
		Conn:db.Conn,
		Help: db.Helper{
			PqArray: pq.Array,
		},
	}
	dl.SimplePingQuery()
	dl.SingleLevelDataQuery()
	dl.SingleLevelWithArrayDataQuery()
	//dl.SingleLevelWithArrayDataInsertQuery()
	//dl.SingleLevelWithArrayDataUpdateQuery()
	dl.SingleLevelWithArrayDataUpdateQueryNoUPDATE()
	dl.SingleLevelWithArrayDataUpdateQueryGetError()

	// Transaction in DB
	dl.SelectInTransaction()
	//dl.InsertInTransaction()
	dl.UpdateInTransaction()

}