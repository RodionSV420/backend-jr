package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//подключаемся к ДБ
func main() {
	fmt.Println("Drivers:", sql.Drivers())
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/v0")
	if err != nil {
		log.Fatal("Unable to open connection to db")
	}
	defer db.Close()
	//АПИ возвращает курсы, которые посещает студент по его ID
	results, err := db.Query("select * from courses")
	if err != nil {
		log.Fatal("Error0", err)
	}
	defer results.Close()
	for results.Next() {
		var (
			coursename string
			Sid        int64
		)
		err = results.Scan(&coursename, &Sid)
		if err != nil {
			log.Fatal("Unable to build course row:", err)
		}
		fmt.Printf("Курс %s посещают студенты с ID %d\n", coursename, Sid)
	}
}
