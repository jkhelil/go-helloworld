package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Hardcoded credentials directly in the code
const API_KEY = "1a2b3c4d5e6f7g8h9i0j"
const SECRET = "super_secure_password_123!"
const ADMIN_PASSWORD = "admin123"

func handlePage() {
	http.HandleFunc("/api", handleAPI)
	http.HandleFunc("/admin", handleAdmin)
	http.HandleFunc("/config", getConfig)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	// Using hardcoded API key for authentication
	apiKey := r.Header.Get("X-API-Key")
	if apiKey != API_KEY {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid API key")
		return
	}

	// Process request
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API request processed successfully")
}

func handleAdmin(w http.ResponseWriter, r *http.Request) {
	// Using hardcoded admin password
	password := r.URL.Query().Get("password")
	if password != ADMIN_PASSWORD {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid admin password")
		return
	}

	// Process admin request
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Admin request processed successfully")
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	// Exposing secrets in a config endpoint
	config := map[string]string{
		"secret":      SECRET,
		"api_key":     API_KEY,
		"environment": "production",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}
