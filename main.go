package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Account struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cpf        int       `json:"cpf"`
	Secret     int       `json:"secret"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at"`
}

var Accounts []Account = []Account{
	{
		Id:         1,
		Name:       "Fulano de Tal",
		Cpf:        11122233344,
		Secret:     552213,
		Balance:    1000,
		Created_at: time.Date(2021, time.April, 16, 12, 0, 0, 0, time.UTC),
	},
	{
		Id:         2,
		Name:       "Siclano de Tal",
		Cpf:        22211133344,
		Secret:     225513,
		Balance:    1500,
		Created_at: time.Date(2021, time.April, 16, 13, 0, 0, 0, time.UTC),
	},
	{
		Id:         3,
		Name:       "Beltrano de Tal",
		Cpf:        33311122244,
		Secret:     135522,
		Balance:    0,
		Created_at: time.Date(2021, time.April, 16, 19, 0, 0, 0, time.UTC),
	},
}

func serverConfig() {
	routesConfig()

	fmt.Println("Servidor está rodando na porta 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func routesConfig() {
	http.HandleFunc("/", root)
	http.HandleFunc("/accounts", accountsList)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá Stone!")
}

func accountsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Accounts)
}

func main() {
	serverConfig()
}
