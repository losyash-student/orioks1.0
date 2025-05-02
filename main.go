package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type App struct {
	psqldb *sql.DB
}

func main() {
	db := pgConn()
	defer db.Close()
	myDb := &App{psqldb: db}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/student", myDb.getItems)
	http.HandleFunc("/student/create", myDb.createItem)
	http.HandleFunc("/student/delete", myDb.deleteItem)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
