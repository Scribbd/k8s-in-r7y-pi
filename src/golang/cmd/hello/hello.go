package main

import (
	"fmt"
	"log"
	"net/http"
)

// HTTP Hello server but cursed
func HewwoServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hewwo, you reqwested: %s", r.URL.Path)
	log.Printf("Weceived weqwest for path: %s", r.URL.Path)
}

func main() {
	var addr string = ":8180"
	handler := http.HandlerFunc(HewwoServer)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Could not listen on port %s %v", addr, err)
	}
}
