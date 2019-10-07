package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {

	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func main() {

	db, err := sql.Open("mysql", "root:1234@/")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "create database if not exists cursoGolang")
	exec(db, "use cursoGolang")
	exec(db, "drop table if exists users")
	exec(db, `create table usuarios(
				id integer auto_increment,
				nome varchar(50),
				PRIMARY KEY(id)
		)`)
}
