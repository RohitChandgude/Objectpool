package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

import (
	"fmt"
)

var database *sql.DB
var acq [10]*sql.DB
var objpool = make(chan *sql.DB)

func Init(n int) {
	objpool = make(chan *sql.DB, n)

	for i := 0; i < n; i++ {
		db, err := sql.Open("mysql","rutuja:rutuja@tcp(127.0.0.1:3306)/go")
		if err != nil {
			fmt.Println("in con error")
		}

		err = db.Ping()
		if err != nil {
			fmt.Println("connection not success")
		} else {
			fmt.Println("connection successful")
		}
		objpool <- db
	}
}

func aquire() *sql.DB {

	obj := <-objpool
	return obj
}

func release(obj *sql.DB) {
	objpool <- obj
}

func main(){
	Init(10)
	fmt.Println(objpool)
	a := aquire()
	b := aquire()
	c := aquire()
	fmt.Println("after acquire\n")
	fmt.Println(objpool)
	fmt.Println(a)
	fmt.Println(c)
	release(b)
	fmt.Println("after release\n")
	fmt.Println(objpool)
}
