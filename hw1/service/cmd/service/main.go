package main

import (
	"net/http"
	"os"
)

import (
	"github.com/RB387/otus-microservices-hw/hw1/service/internal/handlers"
	"log"
)

const defaultAddress = ":8000"

func main() {
	addr := os.Getenv("SERVICE_ADDRESS")
	if addr == "" {
		addr = defaultAddress
	}

	http.HandleFunc("/health/", handlers.HealthHandler)
	http.HandleFunc("/", handlers.NotFoundHandler)

	log.Printf("running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
