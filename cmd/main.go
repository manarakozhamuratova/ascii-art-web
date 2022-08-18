package main

import (
	"ascii/internal/delivery"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", delivery.Home)
	mux.HandleFunc("/result", delivery.Result)

	mux.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./static/style/"))))

	mux.Handle("/favicon/", http.StripPrefix("/favicon/", http.FileServer(http.Dir("./static/favicon/"))))

	log.Println("http://localhost:9090")
	http.ListenAndServe(":9090", mux)
}
