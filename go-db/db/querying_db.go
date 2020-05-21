package db

import (
	"context"
	"database/sql"
	//"database/sql/driver"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"nipun.io/go-cookbook/go-db/domain"
)

//type Helper struct {
//	domain.PqStringArray func(interface{})interface{Scan(interface {}) error; Value() (driver.Value, error)}
//}

type DAO_Layer struct {
	Conn *sqlx.DB
}


func (dl *DAO_Layer) SimplePingQuery() int {
	Conn = dl.Conn
	query := "SELECT 2"
	result := Conn.QueryRow(query)
	var a int
	result.Scan(&a)
	fmt.Println(a)
	return a
}

func (dl *DAO_Layer) SingleLevelDataQuery() []domain.DBUser {
	Conn = dl.Conn
	query := "SELECT id, name FROM db_user;"
	rows , err:= Conn.Query(query)
	if err != nil {
		fmt.Println("Got error : ", err.Error())
	}
	defer rows.Close()
	dbRows := []domain.DBUser{}
	for rows.Next() {
		row := domain.DBUser{}
		err := rows.Scan(&row.Id, &row.Name)
		if err != nil {
			log.Fatal(err)
		}
		dbRows = append(dbRows, row)
		log.Println(row.Id, row.Name)
	}
	return dbRows
}

func (dl *DAO_Layer) SingleLevelWithArrayDataQuery() []domain.DBUser {
	Conn = dl.Conn
	query := "SELECT id, name, contacts FROM db_user;"
	rows , err:= Conn.Query(query)
	if err != nil {
		fmt.Println("Got error : ", err.Error())
	}
	defer rows.Close()
	dbRows := []domain.DBUser{}
	for rows.Next() {
		row := domain.DBUser{}
		err := rows.Scan(&row.Id, &row.Name,&row.Contacts)
		if err != nil {
			log.Fatal(err)
		}
		dbRows = append(dbRows, row)
		log.Println(row.Id, row.Name, row.Contacts)
	}
	return dbRows
}

func (dl *DAO_Layer) SingleLevelWithArrayDataInsertQuery() sql.Result {
	Conn = dl.Conn
	//domain.PqStringArray := dl.Help.domain.PqStringArray
	query := "INSERT INTO db_user VALUES ($1, $2, $3)";
	result, err:= Conn.Exec(query, 3, "user-3", domain.PqStringArray([]string{"1029384756", "0192837465"}))
	if err != nil {
		fmt.Println("Got error : ", err.Error())
	}
	return result
}

func (dl *DAO_Layer) SingleLevelWithArrayDataUpdateQuery() sql.Result {
	Conn = dl.Conn
	//domain.PqStringArray := dl.Help.domain.PqStringArray
	query := "UPDATE db_user SET contacts = $1 WHERE id = $2";
	result , err:= Conn.Exec(query, domain.PqStringArray([]string{"111111111", "000000000"}), 3)
	if err != nil {
		fmt.Println("Got error : ", err.Error())
	}
	return result
}

func (dl *DAO_Layer) SingleLevelWithArrayDataUpdateQueryNoUPDATE() sql.Result {
	Conn = dl.Conn
	//domain.PqStringArray := dl.Help.domain.PqStringArray
	query := "UPDATE db_user SET contacts = $1 WHERE id = $2";
	result , err:= Conn.Exec(query, domain.PqStringArray([]string{"1111111111", "000000000"}), 10)
	if err != nil {
		fmt.Println("Got error : ", err.Error())
	}
	fmt.Println(result.RowsAffected())
	return result
}

func (dl *DAO_Layer) SingleLevelWithArrayDataUpdateQueryGetError() sql.Result {
	Conn = dl.Conn
	query := "UPDATE db_user SET contacts = $1 WHERE id = $2";
	result , err:= Conn.Exec(query,"1111111111", 3)
	if err != nil {
		fmt.Println("Got error : ", err.Error())
	}

	return result
}

func (dl *DAO_Layer) SelectInTransaction() sql.Result {
	Conn = dl.Conn
	tx, err := Conn.BeginTxx(context.Background(), nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	result, err := tx.Exec("SELECT * FROM db_user where id=1;")
	if err != nil {
		tx.Rollback()
	}
	rowsEffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("rows effected : ", rowsEffected)
	err = tx.Commit()
	if err != nil {
		fmt.Println(err.Error())
	}

	return result
}

func (dl *DAO_Layer) InsertInTransaction() sql.Result {
	Conn = dl.Conn
	//domain.PqStringArray := dl.Help.domain.PqStringArray
	tx, err := Conn.BeginTxx(context.Background(), nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	query := "INSERT INTO db_user VALUES ($1, $2, $3)";
	result , err:= tx.Exec(query, 4, "user-4", domain.PqStringArray([]string{"2222222222", "3333333333"}))
	if err != nil {
		fmt.Println("Got error : ", err.Error())
		tx.Rollback()
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}

	rowsEffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Rows effected while inserting :", rowsEffected)
	return result
}

func (dl *DAO_Layer) UpdateInTransaction() sql.Result {
	Conn = dl.Conn
	//domain.PqStringArray := dl.Help.domain.PqStringArray
	tx, err := Conn.BeginTxx(context.Background(), nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	query := "UPDATE db_user SET contacts = $1 WHERE id = $2";
	result , err:= tx.Exec(query, domain.PqStringArray([]string{"4454444444", "555555555"}), 4)
	if err != nil {
		fmt.Println("Got error : ", err.Error())
		tx.Rollback()
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		fmt.Print("commiting again : ", err)
	}

	rowsEffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Rows effected while updating :", rowsEffected)
	return result
}


