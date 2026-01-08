package api

import (
	"encoding/json"
	"net/http"
)

var apps = make(map[string]App)

func DeployApp(w http.ResponseWriter, r *http.Request) {
	var app App
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	app.Status = "pending"
	apps[app.Name] = app

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(app)
}

func ListApps(w http.ResponseWriter, _ *http.Request) {
	var list []App
	for _, app := range apps {
		list = append(list, app)
	}
	json.NewEncoder(w).Encode(list)
}
