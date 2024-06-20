package main

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	Port string
}

func main() {
	server := App{
		Port: "8080",
	}
	server.Start()
}

func (a App) Start() {
	setUpRoutes()

	addr := fmt.Sprintf(":%s", a.Port)
	log.Printf("Starting app at localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func setUpRoutes() {
	http.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong\n")
	}))
}
