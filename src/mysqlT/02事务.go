package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:123@tcp(localhost:3306)/study")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	_, err1 := tx.Exec("insert into person(name,age,rmb,gender,birthday)values(?,?,?,?,?);", "小猪", 22, 50.09, false, 19971008)
	_, err2 := tx.Exec("insert into person(name,age,rmb,gender,birthday)values(?,?,?,?,?);", "ztt", 22, 50.09, false, 19971008)
	if err1 != nil || err2 != nil {
		panic(err)
		tx.Rollback()
	} else {
		tx.Commit()
	}

}
