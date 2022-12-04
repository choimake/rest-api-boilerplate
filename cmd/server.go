package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rest-api-boilerplate/internal/driver/handler/chi"
	"rest-api-boilerplate/internal/driver/registry"
)

func main() {

	r := registry.Initialize()

	// DI
	registry.InitProvider(r)
	p := registry.GetProvider()

	// router
	mux := chi.NewRouter(p)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, mux)
	fmt.Printf("[START] server. port: %s\n", port)
	if err != nil {
		log.Fatal(err)
	}
}
