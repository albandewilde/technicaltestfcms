package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", isUp)

	log.Println("Server started on 0.0.0.0:8254")
	log.Fatal(http.ListenAndServe("0.0.0.0:8254", nil))
}

func isUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is up and running")
}
