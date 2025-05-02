package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "db" //172.17.0.2
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "persons"
)

func delStudent(db *sql.DB, id int) {
	psqlInfo := fmt.Sprintf("call delete_student_by_id(%d)", id)
	db.Exec(psqlInfo)
}

func insertStudent(db *sql.DB, id int, name string, surname string, mark int) {
	sqlStatement := `
	INSERT INTO student (id, name, surname, mark)
	VALUES ($1, $2, $3, $4)`
	db.Exec(sqlStatement, id, name, surname, mark)
}

func getAllStudents(db *sql.DB) []Student {
	rows, err := db.Query("select * from student")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	persons := []Student{}
	for rows.Next() {
		p := Student{}
		err := rows.Scan(&p.ID, &p.Surname, &p.Name, &p.Mark)
		if err != nil {
			fmt.Println(err)
			continue
		}
		persons = append(persons, p)
	}
	return persons
}

func printStudents(db *sql.DB) {
	rows, err := db.Query("select * from student")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []Student{}

	for rows.Next() {
		p := Student{}
		err := rows.Scan(&p.ID, &p.Surname, &p.Name, &p.Mark)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.ID, p.Mark, p.Name, p.Surname)
	}
}

func pgConn() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var db *sql.DB
	var err error
	fmt.Println(psqlInfo)
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("Successfully connected!")
				return db
			}
			fmt.Println(err)
		}

		time.Sleep(5_000_000_000)
	}
	fmt.Println("Error!")
	return nil
}
