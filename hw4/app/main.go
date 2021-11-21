package main

import (
	"fmt"
	"log"
	"net/http"
)

const addr = ":8080"

func versionHandler(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("{\"version\": \"v2\"}")); err != nil {
		fmt.Println("error in health handler: ", err)
		w.WriteHeader(500)
	}

	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/version/", versionHandler)

	log.Printf("running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
