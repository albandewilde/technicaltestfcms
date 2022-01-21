package handlers

import (
	"fmt"
	"net/http"
)

func Alerts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.HandlerFunc(get).ServeHTTP(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Consult your alerts")
}
