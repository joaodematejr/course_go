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
	stmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
	stmt.Exec("Uoxiton Istive", 1)
	stmt.Exec("Sheristone Uoxiton", 2)

	stmt2, _ := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	stmt2.Exec(3)
}
