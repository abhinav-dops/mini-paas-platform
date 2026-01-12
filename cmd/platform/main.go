package main

import (
	"log"
	"net/http"

	"github.com/abhinav-dops/mini-paas-platform/internal/api"
)

func main() {
	log.Println("Mini PaaS Platform starting...")

	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/apps/deploy", api.DeployApp)
	http.HandleFunc("/apps", api.ListApps)
	http.HandleFunc("/infra/provision", api.ProvisionInfra)
	http.HandleFunc("/infra/status", api.GetInfraStatus)
	http.HandleFunc("/infra/destroy", api.DestroyInfra)

	log.Fatal(http.ListenAndServe(":9000", nil))
}

// atleast wait 1min after provisioning infra before deploying apps
