package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	users := make(map[string]string)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		address := r.FormValue("address")

		if name == "" || address == "" {
			http.Error(w, "Name and address are required", http.StatusBadRequest)
			return
		}

		users[name] = address

		fmt.Fprintf(w, "User %s has been created successfully.", name)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
