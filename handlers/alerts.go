package handlers

import (
	"fmt"
	"net/http"

	"github.com/albandewilde/technicaltestfcms/alerts"
)

type AlertsHandlers struct {
	alerts *alerts.Alerts
}

func NewAlertsHandlers(as *alerts.Alerts) AlertsHandlers {
	return AlertsHandlers{alerts: as}
}

func (ah *AlertsHandlers) Alerts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.HandlerFunc(ah.list).ServeHTTP(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (ah *AlertsHandlers) list(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Consult your alerts")
}

func (ah *AlertsHandlers) Alert(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		http.HandlerFunc(ah.addAlert).ServeHTTP(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (ah *AlertsHandlers) addAlert(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You create a new alert")
}
