package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Natannms/go_sqlc/infra/service/db"
)

func Home(r http.ResponseWriter, rq *http.Request) {
	json.NewEncoder(r).Encode("HOME")
}

func CreateUser(r http.ResponseWriter, rq *http.Request) {
	// var user models.User
	// json.NewDecoder(rq.Body).Decode(&user)
	// database.Conn().Create(&user)
	// json.NewEncoder(r).Encode(user)


	dbcon, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sqlc")
	if err != nil {
		panic(err.Error())
	}

	defer dbcon.Close()

	//execute query
	query := db.New(dbcon)

	ctx := context.Background()
	users, err := query.CreateUser(ctx)
	if err != nil {
		json.NewEncoder(r).Encode("Falha ao executar consulta getUsers")

	}

	json.NewEncoder(r).Encode(users)
}

func AllUsers(r http.ResponseWriter, rq *http.Request) {

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
		json.NewEncoder(r).Encode("Falha ao executar consulta getUsers")

	}

	json.NewEncoder(r).Encode(users)
}

// func GetUser(r http.ResponseWriter, rq *http.Request) {
// 	vars := mux.Vars(rq)
// 	id := vars["id"]
// 	var user models.User

// 	result := database.Conn().First(&user, id)
// 	if result.Error != nil {
// 		http.Error(r, result.Error.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(r).Encode(user)
// }

// func UpdateUser(r http.ResponseWriter, rq *http.Request) {
// 	vars := mux.Vars(rq)
// 	id := vars["id"]

// 	var user models.User

// 	result := database.Conn().First(&user, id)
// 	if result.Error != nil {
// 		http.Error(r, result.Error.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewDecoder(rq.Body).Decode(&user)
// 	database.Conn().Save(&user)
// 	json.NewEncoder(r).Encode(&user)
// }

// func DeleteUser(r http.ResponseWriter, rq *http.Request) {
// 	vars := mux.Vars(rq)
// 	id := vars["id"]

// 	var user models.User

// 	result := database.Conn().First(&user, id)
// 	if result.Error != nil {
// 		http.Error(r, result.Error.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	database.Conn().Delete(&user, id)
// 	json.NewEncoder(r).Encode(&user)
// }
