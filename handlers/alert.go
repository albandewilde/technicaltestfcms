package handlers

import (
	"fmt"
	"net/http"
)

func Alert(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "handle alert")
}
