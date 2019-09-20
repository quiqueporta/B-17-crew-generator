package main

import (
	"log"
	"net/http"

	"github.com/quiqueporta/b17-crew-generator/pdf"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/pdf")
	w.Write(pdf.GeneratePDF())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
