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
	// This is our "Digital Filing Cabinet" (A tiny database)
	// It stores the Bag ID and the Passenger's Name
	airportDatabase := map[string]string{
		"BAG111": "Dilukshi Wijayawardhana (Flight UL101)",
		"BAG222": "John Doe (Flight UL202)",
	}

	// 1. Show the visual dashboard
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// 2. The DYNAMIC Scan API
	http.HandleFunc("/api/scan", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		// The code reads the Bag ID from the scanner (the button click)
		scannedID := r.URL.Query().Get("id")

		// The code checks the Digital Filing Cabinet
		passengerInfo, bagExists := airportDatabase[scannedID]

		var dynamicMessage string

		if bagExists {
			// If the bag is in the database, say who it belongs to!
			dynamicMessage = fmt.Sprintf("✅ SUCCESS: Bag %s loaded! Owner: %s", scannedID, passengerInfo)
		} else {
			// If the bag is NOT in the database, sound an alarm!
			dynamicMessage = fmt.Sprintf("❌ SECURITY ALERT: Bag %s does not belong to any passenger!", scannedID)
		}

		// Send the dynamic message back to the screen
		json.NewEncoder(w).Encode(SystemResponse{Message: dynamicMessage})
	})

	fmt.Println("AirTrack Global is LIVE! Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
