package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/abhinav-dops/mini-paas-platform/internal/config"
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

	if app.ContainerPort == 0 {
		http.Error(w, "container_port is required", http.StatusBadRequest)
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

	image := fmt.Sprintf("mini-paas%s:latest", app.Name)

	if err := docker.CloneRepo(Infra.IP, config.SSHKeyPath, app.Repo); err != nil {
		app.Status = "failed"
		app.Error = err.Error()
		apps[appName] = app
		return
	}

	if err := docker.BuildImage(Infra.IP, config.SSHKeyPath, image); err != nil {
		app.Status = "failed"
		app.Error = err.Error()
		apps[appName] = app
		return
	}

	err := docker.RunRemoteContainer(
		Infra.IP,
		config.SSHKeyPath,
		app.Name,
		app.Port,
		app.ContainerPort,
		image,
	)

	if err := docker.HealthCheck(Infra.IP, config.SSHKeyPath, app.Port); err != nil {
		app.Status = "unhealthy"
		app.Error = "health check failed"
		return
	}

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

func DestroyApp(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/apps/")

	if _, ok := apps[name]; !ok {
		http.Error(w, "app not found", 404)
		return
	}

	err := docker.RemoveContainer(
		Infra.IP,
		config.SSHKeyPath,
		name,
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	delete(apps, name)
	w.Write([]byte("deleted"))
}
