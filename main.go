package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/albandewilde/technicaltestfcms/alerts"
	"github.com/albandewilde/technicaltestfcms/handlers"
)

var ah handlers.AlertsHandlers

func main() {
	http.HandleFunc("/", isUp)

	http.HandleFunc("/alert/", ah.Alert)
	http.HandleFunc("/alerts/", ah.Alerts)

	log.Println("Server started on 0.0.0.0:8254")
	log.Fatal(http.ListenAndServe("0.0.0.0:8254", nil))
}

func init() {
	a0, _ := alerts.NewAlert("a", "bob", "0001-01-01", "alrt1", 0, 10, "Nantes")
	a1, _ := alerts.NewAlert("b", "bob", "0001-01-01", "alrt2", 1, 50, "Orléans")
	a2, _ := alerts.NewAlert("c", "alice", "0001-01-01", "my_alerte", 100, 1000, "Paris")
	a3, _ := alerts.NewAlert("d", "alice", "0001-01-01", "my second alerte", 5, 5, "Rennes")
	a4, _ := alerts.NewAlert("e", "alice", "0001-01-01", "my third", 4, 9, "Lyon")

	as := alerts.NewAlerts(
		map[string][]alerts.Alert{
			"bob":   {a0, a1},
			"alice": {a2, a3, a4},
		},
	)

	ah = handlers.NewAlertsHandlers(&as)
}

func isUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is up and running")
}
