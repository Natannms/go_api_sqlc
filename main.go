package main

import (
	"fmt"

	"github.com/Natannms/go_sqlc/interface/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Running in port http://localhost:8000") // Mensagem com link para executar no nvegador
	routes.HandleRequest()
}
