package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Bullshit little echo test from POD-1 to POD-2
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "first pod")
	})
	http.ListenAndServe(":3000", r)
}
