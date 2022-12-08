package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:go242526@/curso_go")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO usuarios(id, nome) VALUES(? ?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")

	_, err = stmt.Exec(1, "Ana") // chave duplicada

	if err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()

}
