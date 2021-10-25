package handlers

import (
	"fmt"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got error 404 on ", r.RequestURI)

	_, err := w.Write([]byte(fmt.Sprintf("404, path %s not found", r.RequestURI)))
	if err != nil {
		fmt.Print("error in 404 handler ", err)
	}
}
