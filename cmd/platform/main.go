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

	log.Fatal(http.ListenAndServe(":9000", nil))
}
