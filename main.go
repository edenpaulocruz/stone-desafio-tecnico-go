package main

import (
	"fmt"
	"net/http"
)

func serverConfig() {
	routesConfig()

	fmt.Println("Servidor está rodando na porta 8888")
	http.ListenAndServe(":8888", nil)
}

func routesConfig() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá Stone!")
}

func main() {
	serverConfig()
}
