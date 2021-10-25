package handlers

import (
	"fmt"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("{\"status\": \"OK\"}")); err != nil {
		fmt.Println("error in health handler: ", err)
		w.WriteHeader(500)
	}

	w.WriteHeader(200)
}
