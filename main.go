package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

// GetBookingDetails - Get all booking details using the stored procedure
func GetBookingDetails(db *sql.DB, professionalId, partnerId, bookingId, langId *int) (*BookingDetailsResult, error) {
	result := &BookingDetailsResult{}

	// Execute the stored procedure
	rows, err := db.Query("EXEC dbo.getBookingDetails @p1, @p2, @p3, @p4",
		professionalId, partnerId, bookingId, langId)
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %v", err)
	}
	defer rows.Close()

	// Scan all recordsets
	err = scanMultipleRecordsets(rows, result)
	if err != nil {
		return nil, fmt.Errorf("error scanning recordsets: %v", err)
	}

	return result, nil
}

func main() {
	connString := "server=localhost;user id=user_node_boosterdb;password=6nWNlTX^LpGRONb$WTmeH;database=boosterdb"

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	// Setup HTTP routes
	http.HandleFunc("/booking/", func(w http.ResponseWriter, r *http.Request) {
		handleBookingDetails(w, r, db)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleHome(w, r)
	})

	// Start server
	fmt.Println("Server starting on :8080...")
	fmt.Println("Access booking details: http://localhost:8080/booking/1404590")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func logBookingDetailsAsJSON(result *BookingDetailsResult) {
	// Convert to JSON with pretty printing
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Printf("Error marshaling to JSON: %v", err)
		return
	}

	fmt.Println("=== COMPLETE BOOKING DETAILS (JSON) ===")
	fmt.Println(string(jsonData))

	// Also log summary counts
	fmt.Printf("\n=== SUMMARY ===\n")
	fmt.Printf("Booking Details: %d records\n", len(result.BookingDetails))
	fmt.Printf("Hotel Details: %d records\n", len(result.HotelDetails))
	fmt.Printf("Detail Options: %d records\n", len(result.DetailOptions))
	fmt.Printf("Booking Services: %d records\n", len(result.BookingServices))
	fmt.Printf("Miscellaneous: %d records\n", len(result.Miscellaneous))
	fmt.Printf("Passengers: %d records\n", len(result.Passengers))
	fmt.Printf("Booking Summary: %d records\n", len(result.BookingSummary))
	fmt.Printf("Booking Hotels: %d records\n", len(result.BookingHotels))
	fmt.Printf("Booking Products: %d records\n", len(result.BookingProducts))
	fmt.Printf("Transactions: %d records\n", len(result.Transactions))
	fmt.Printf("Cancellation Policies: %d records\n", len(result.CancellationPolicies))
	fmt.Printf("Booking Logs: %d records\n", len(result.BookingLogs))
}

// HTTP Handlers

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Booking Details API</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        h1 { color: #333; }
        .example { background: #f0f0f0; padding: 15px; border-radius: 5px; margin: 20px 0; }
        .endpoint { color: #0066cc; font-weight: bold; }
    </style>
</head>
<body>
    <h1>Booking Details API</h1>
    <p>Available endpoints:</p>
    
    <div class="example">
        <h3>Get Booking Details</h3>
        <p><span class="endpoint">GET /booking/{bookingId}</span></p>
        <p>Returns complete booking details in JSON format</p>
        <p>Example: <a href="/booking/1404590">/booking/1404590</a></p>
    </div>

    <div class="example">
        <h3>Query Parameters (optional)</h3>
        <p><span class="endpoint">GET /booking/{bookingId}?professionalId=1&partnerId=2&langId=1</span></p>
        <p>All parameters are optional and will be set to NULL if not provided</p>
    </div>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func handleBookingDetails(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extract booking ID from URL path
	path := r.URL.Path
	if len(path) <= 9 { // "/booking/" is 9 characters
		http.Error(w, "Booking ID is required", http.StatusBadRequest)
		return
	}

	bookingIdStr := path[9:] // Extract everything after "/booking/"
	bookingId, err := strconv.Atoi(bookingIdStr)
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	// Get optional query parameters
	var professionalId, partnerId, langId *int

	if profId := r.URL.Query().Get("professionalId"); profId != "" {
		if id, err := strconv.Atoi(profId); err == nil {
			professionalId = &id
		}
	}

	if partId := r.URL.Query().Get("partnerId"); partId != "" {
		if id, err := strconv.Atoi(partId); err == nil {
			partnerId = &id
		}
	}

	if lang := r.URL.Query().Get("langId"); lang != "" {
		if id, err := strconv.Atoi(lang); err == nil {
			langId = &id
		}
	}

	// Get booking details
	result, err := GetBookingDetails(db, professionalId, partnerId, &bookingId, langId)
	if err != nil {
		log.Printf("Error getting booking details: %v", err)
		http.Error(w, "Error retrieving booking details", http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	// Also log to console
	logBookingDetailsAsJSON(result)
}
