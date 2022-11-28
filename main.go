package main

import (
	"log"
	"net/http"

	"github.com/MrHenri/go-expert-chalange-1/server"
)

func main() {
	http.HandleFunc("/cotacao", server.Quotation)
	println("Server started on Port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
