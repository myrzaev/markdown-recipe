package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
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
	f, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		slog.Error("failed to get index.html", "filename", "index.html")
		return
	}
	w.Write(f)
}
