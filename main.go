package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", defaultOne)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func defaultOne(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Shrey!"))
}
