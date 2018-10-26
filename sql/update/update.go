package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:senha@/cursogo")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	res, _ := stmt.Exec("Maria bla", 2)
	linhas, _ := res.RowsAffected()

	fmt.Println(linhas)
}
