package main

import (
	"encoding/json"
	"net/http"
)

var Students = []Student{}

func (app *App) getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Students = getAllStudents(app.psqldb)
	json.NewEncoder(w).Encode(Students)
}

func (app *App) createItem(w http.ResponseWriter, r *http.Request) {
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	insertStudent(app.psqldb, student.ID, student.Name, student.Surname, student.Mark)
}

func (app *App) deleteItem(w http.ResponseWriter, r *http.Request) {
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	delStudent(app.psqldb, student.ID)
}
