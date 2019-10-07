package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:1234@/cursoGolang")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id,nome) values(?,?)")

	stmt.Exec(1000, "Teste")
	stmt.Exec(2000, "Teste2")
	stmt.Exec(1, "Teste3")

	_, err = stmt.Exec(1, "Teste4")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}
