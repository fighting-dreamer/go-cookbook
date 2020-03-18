package db

import "fmt"

func SimplePingQuery() {
	query := "SELECT 2"
	result := Conn.QueryRow(query)
	var a int
	result.Scan(&a)
	fmt.Println(a)
}

