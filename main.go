package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Natannms/go_sqlc/infra/service/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbcon, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sqlc")
	if err != nil {
		panic(err.Error())
	}

	defer dbcon.Close()

	//execute query
	query := db.New(dbcon)

	ctx := context.Background()
	users, err := query.GetUsers(ctx)
	if err != nil {
		log.Fatalf("Falha ao executar consulta GetUsers: %v", err)
	}

	for _, u := range users {
		log.Printf("Usuario: %v", u)
	}

	fmt.Println(users)

}
