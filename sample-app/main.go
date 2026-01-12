package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello from mini paas")
	})

	fmt.Println("server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
