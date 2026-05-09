package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type SystemResponse struct {
	Message string `json:"message"`
}

func main() {
	// In-memory datastore mapping Bag IDs to passenger records
	airportDatabase := map[string]string{
		"BAG111": "Dilukshi Wijayawardhana (Flight UL101)",
		"BAG222": "John Doe (Flight UL202)",
	}

	// Route: Serve the frontend UI dashboard
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Route: Process incoming bag scans dynamically
	http.HandleFunc("/api/scan", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		// Extract scanned ID from URL query parameters
		scannedID := r.URL.Query().Get("id")

		// Lookup bag record in datastore
		passengerInfo, bagExists := airportDatabase[scannedID]

		var dynamicMessage string

		if bagExists {
			// Match found: Construct success payload
			dynamicMessage = fmt.Sprintf("SUCCESS: Bag %s loaded! Owner: %s", scannedID, passengerInfo)
		} else {
			// No match: Trigger security alert payload
			dynamicMessage = fmt.Sprintf("SECURITY ALERT: Unregistered Bag %s detected!", scannedID)
		}

		// Transmit JSON response back to the client
		json.NewEncoder(w).Encode(SystemResponse{Message: dynamicMessage})
	})

	// Initialize server on port 8080
	fmt.Println("AirTrack API is LIVE! Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}