package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	remindHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "reminder endpoint\n")
	}

	http.HandleFunc("/remind", remindHandler)
		log.Println("Listing for requests at http://localhost:8000/remind")
	
	
	cpuHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "task manager endpoint\n")
	}

	http.HandleFunc("/cpu", cpuHandler)
		log.Println("Listing for requests at http://localhost:8000/cpu")

	log.Fatal(http.ListenAndServe(":8000", nil))
}