package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	http.HandleFunc("/accounts", accountsRoutes)
	http.HandleFunc("/accounts/", accountsRoutes)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá Stone!")
}

func accountsRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParts := strings.Split(r.URL.Path, "/")

	if len(urlParts) == 2 || len(urlParts) == 3 && urlParts[2] == "" {
		switch r.Method {
		case "GET":
			accountsList(w, r)
		case "POST":
			accountsAdd(w, r)
		}
	} else {
		w.WriteHeader((http.StatusNotFound))
	}
}

func accountsList(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(Accounts)
}

func accountsAdd(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	// Tratar erro

	var newAccount Account
	json.Unmarshal(body, &newAccount)
	newAccount.Id = Accounts[len(Accounts)-1].Id + 1
	newAccount.Created_at = time.Now()
	Accounts = append(Accounts, newAccount)

	encoder := json.NewEncoder(w)
	encoder.Encode(newAccount)
}

func main() {
	serverConfig()
}
