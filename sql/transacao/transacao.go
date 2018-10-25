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

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?, ?)")
	stmt.Exec(2001, "Tiago")
	stmt.Exec(2002, "Jose")
	_, err = stmt.Exec(10, "Pedro")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
	fmt.Println("commit")
}
