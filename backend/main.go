package main

import (
	"challenge/config"
	"challenge/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	port := config.GetPort()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/contact", handlers.HandleContactForm)

	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
