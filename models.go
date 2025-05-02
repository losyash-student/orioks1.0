package main

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Mark    int    `json:"mark"`
}

//CREATE TABLE student (id integer, name varchar(255), surname varchar(255), mark int);
