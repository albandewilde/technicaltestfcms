package handlers

import (
	"fmt"
	"net/http"
)

func Alert(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		http.HandlerFunc(post).ServeHTTP(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You create a new alert")
}
