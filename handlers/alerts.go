package handlers

import (
	"fmt"
	"net/http"
)

func Alerts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "handle alerts")
}
