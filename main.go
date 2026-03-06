package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	ID   int
	Nome string
}

func users(w http.ResponseWriter, r *http.Request) {

	file, _ := os.Open("users.json")

	var lista []User

	json.NewDecoder(file).Decode(&lista)

	file.Close()

	json.NewEncoder(w).Encode(lista)
}

func main() {

	http.HandleFunc("/users", users)

	fmt.Println("Servidor rodando em http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
