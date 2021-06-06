package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)


func FetchNameById(db *sql.DB, id int) (string, error) {
	username := ""
	err := db.QueryRow("SELECT username FROM users WHERE id=?", id).Scan(&username)
	return username, errors.Wrap(err, "main.FetchNameByID Fetch No Result")
}

func main() {
	// Fake DB
	db := sql.DB{}
	
	username, err := FetchNameById(&db, 10)
	// 在调用Dao层方法后上层对返回结果进行统一区分
	if err != nil{
		if errors.Cause(err) == sql.ErrNoRows {
			fmt.Printf("Data Not Found, %v \n", err)
			fmt.Printf("%+v \n", err)
			return
		}
		// Deal with other errors
	}
	fmt.Print(username)
}
