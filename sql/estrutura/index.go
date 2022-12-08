package main

import (
	"database/sql"
	"fmt"

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
	db, err := sql.Open("mysql", "root:go242526@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	exec(db, "CREATE DATABASE IF NOT EXISTS curso_go")
	exec(db, "USE curso_go")
	exec(db, "DROP TABLE IF EXISTS usuarios")
	exec(db, `CREATE TABLE usuarios (
		id INTEGER AUTO_INCREMENT,
		nome VARCHAR(80),
		PRIMARY KEY (id)
	)`)

	fmt.Println("Tabela criada com sucesso!")

}
