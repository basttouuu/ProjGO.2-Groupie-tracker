package main

import (
	"API_HP/src/routers"
	"API_HP/src/templates"
	"fmt"
	"log"
	"net/http"
)

func main() {
	templates.Load()

	mux := routers.MainRouter()

	addr := "localhost:8080"
	fmt.Printf("Serveur prêt sur http://%s\n", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
