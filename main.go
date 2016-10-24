package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8090", http.HandlerFunc(requestHandler)); err != nil {
		log.Fatalf("error in server: %s", err)
	}
}

func requestHandler(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/json":
		jsonHandler(w)
	default:
		http.Error(w, "page not found", http.StatusNotFound)
	}
}

func jsonHandler(w http.ResponseWriter) {
	w.Header().Set("ContentType", "application/json; charset=utf-8")

	a := make(map[string]int, 10)
	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("Test Number %d", i+1)
		a[k] = rand.Intn(1000)
	}
	if err := json.NewEncoder(w).Encode(a); err != nil {
		log.Printf("error when encoding json: %s", err)
	}
}
