package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Mini PaaS Platform starting...")

	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}
