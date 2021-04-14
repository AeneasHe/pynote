package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	DB *sql.DB
}

func NewDataBase(filename string, flag string) DataBase {
	var db = DataBase{}
	// 存在时，先删除再创建
	if flag == "-" {
		os.Remove(filename)
	}
	_db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}
	db.DB = _db
	return db
}

func (db *DataBase) Close() {
	db.DB.Close()
}

func (db *DataBase) CreateTable() {
	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err := db.DB.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func (db *DataBase) Transaction() {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

func (db *DataBase) Query() {
	rows, err := db.DB.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DataBase) TransactionQuery() {
	stmt, err := db.DB.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

func (db *DataBase) Delete() {
	_, err := db.DB.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DataBase) Insert() {
	_, err := db.DB.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DataBase) query() {
	rows, err := db.DB.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
