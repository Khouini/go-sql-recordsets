package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

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

	// Get booking details
	bookingId := 1404590

	result, err := GetBookingDetails(db, nil, nil, &bookingId, nil)
	if err != nil {
		log.Fatal("Error getting booking details:", err)
	}

	// Log complete data as JSON
	logBookingDetailsAsJSON(result)
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
