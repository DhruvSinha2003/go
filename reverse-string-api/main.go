package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ReverseRequest struct {
	InputString string `json:"input"`
}

type ReverseResponse struct {
	Original string `json:"original"`
	Reversed string `json:"reversed"`
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]

	}
	return string(runes)
}

func reverseHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow requests from any origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight OPTIONS request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only accept POST requests
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var req ReverseRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Generate response
	resp := ReverseResponse{
		Original: req.InputString,
		Reversed: reverseString(req.InputString),
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/api/reverse", reverseHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	port := "5050"
	fmt.Printf("Server started at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
