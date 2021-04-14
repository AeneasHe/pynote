package main

import "pynote/db"

func main() {
	db := db.NewDataBase("./foo.db", "-")
	db.CreateTable()
	db.Insert()
	db.Query()
}
