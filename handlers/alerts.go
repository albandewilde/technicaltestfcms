package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	// Check if the body isn't nil
	if r.Body == nil {
		http.Error(w, "No body given", http.StatusBadRequest)
		return
	}

	// Read the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading the body", http.StatusInternalServerError)
		return
	}

	// Parse username
	var uname string
	err = json.Unmarshal(body, &uname)

	if uname == "" {
		http.Error(w, "No name in the body", http.StatusBadRequest)
		return
	}

	// Fetch user name alerts
	as, err := ah.alerts.List(uname)
	if errors.Is(alerts.ErrUserNotFound, err) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Send back username alerts
	jsn, err := json.Marshal(as)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsn)
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
