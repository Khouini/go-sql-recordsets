package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// GetTestData - Simple method that returns all data in one struct
func GetTestData(db *sql.DB) (*TestDataResult, error) {
	result := &TestDataResult{}

	rows, err := db.Query("EXEC sp_GetTestData")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// One line to scan all recordsets
	err = scanMultipleRecordsets(rows, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func main() {
	connString := "server=localhost,6000;user id=sa;password=95W3x4P4kki5;database=TestDB_X"

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Call single recordset procedure
	customers, err := GetCustomers(db)
	if err != nil {
		log.Fatal("Error getting customers:", err)
	}

	fmt.Println("=== CUSTOMERS FROM sp_GetCustomers ===")
	for _, customer := range customers {
		// fmt.Printf("ID: %d, Name: %s, Country: %s\n",
		// 	customer.CustomerID, customer.Name, customer.Country)
		fmt.Printf("ID: %d, Name: %s\n",
			customer.CustomerID, customer.Name)
	}

	// Call multiple recordset procedure
	result, err := GetTestData(db)
	if err != nil {
		log.Fatal("Error getting test data:", err)
	}

	// Print results
	fmt.Println("=== CUSTOMERS ===")
	for _, customer := range result.Customers {
		// fmt.Printf("ID: %d, Name: %s, Country: %s\n",
		// 	customer.CustomerID, customer.Name, customer.Country)
		fmt.Printf("ID: %d, Name: %s\n",
			customer.CustomerID, customer.Name)
	}

	fmt.Println("\n=== ORDERS ===")
	for _, order := range result.Orders {
		fmt.Printf("ID: %d, Customer: %d, Date: %s, Total: %.2f\n",
			order.OrderID, order.CustomerID, order.OrderDate, order.Total)
	}

	fmt.Println("\n=== PRODUCTS ===")
	for _, product := range result.ProductsEmpty {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n",
			product.ProductID, product.Name, product.Price)
	}

	fmt.Println("\n=== PRODUCTS ===")
	for _, product := range result.Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n",
			product.ProductID, product.Name, product.Price)
	}
}

// GetCustomers - Simple method for single recordset procedure
func GetCustomers(db *sql.DB) ([]Customer, error) {
	var customers []Customer

	rows, err := db.Query("EXEC sp_GetCustomers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan single recordset
	err = scanRecordset(rows, &customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}
