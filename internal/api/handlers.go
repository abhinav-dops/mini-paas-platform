package api

import (
	"encoding/json"
	"net/http"

	"github.com/abhinav-dops/mini-paas-platform/internal/docker"
)

var apps = make(map[string]App)

func DeployApp(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if Infra.Status != "ready" {
		http.Error(w, "infra not ready", http.StatusPreconditionFailed)
		return
	}

	var app App
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	app.Status = "pending"
	apps[app.Name] = app

	// async execution
	go executeDeployment(app.Name)

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

func executeDeployment(appName string) {
	app := apps[appName]

	if Infra.Status != "ready" {
		app.Status = "failed"
		app.Error = "infra not ready"
		apps[appName] = app
		return
	}

	err := docker.RunRemoteContainer(
		Infra.IP,
		"C:/Users/abhin/.ssh/mini-paas-key.pem",
		app.Name,
		app.Port,
		"sample:latest",
	)

	if err != nil {
		app.Status = "failed"
		app.Error = err.Error()
		apps[appName] = app
		return
	}

	app.Status = "running"
	apps[appName] = app
	// log.Printf("app %s running", app.Name)
}
