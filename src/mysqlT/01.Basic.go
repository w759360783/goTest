package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	Name  string  `db:"name"`
	Age   int     `db:"age"`
	Money float64 `db:"rmb"`
}

func main001() {
	db, err := sqlx.Open("mysql", "root:123@tcp(localhost:3306)/study")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("Insert into person(name,age,rmb,gender,birthday)values(?,?,?,?,?);", "吴铭", 23, 45.23, true, 19961008)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	affected, err := result.RowsAffected()
	fmt.Println("insertId", insertId)
	fmt.Println("affected", affected)

}

func main (){
	db, err := sqlx.Open("mysql", "root:123@tcp(localhost:3306)/study")
	if err !=nil{
		panic(err)
	}
	var person []Person

	err = db.Select(&person, "select name,age,rmb from person;")

	if err != nil {
		panic(err)
	}
	fmt.Println("查询的结果为：",person)
}
