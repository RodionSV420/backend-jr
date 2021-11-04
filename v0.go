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
	//АПИ строит отчёт для военкомата, путём вывода всех студентов мальчиков, которым сейчас 18 лет. На вход АПИ принимает два параметра возраст студента и пол.
	results, err := db.Query("select * from student where age = 18 and sex = 'M'")
	if err != nil {
		log.Fatal("Error0", err)
	}
	defer results.Close()
	for results.Next() {
		var (
			ID        int64
			full_name string
			age       int64
			sex       string
		)
		err = results.Scan(&ID, &full_name, &age, &sex)
		if err != nil {
			log.Fatal("Unable to build student row:", err)
		}
		fmt.Printf("%sу ровно %d лет и пол %sужской\n", full_name, age, sex)
	}
}
