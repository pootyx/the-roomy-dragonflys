package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)

	r := mux.NewRouter()
	r.HandleFunc("/", HelloServer)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
