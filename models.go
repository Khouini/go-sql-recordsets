package main

type Order struct {
	OrderID    int     `db:"OrderID"`
	CustomerID int     `db:"CustomerID"`
	OrderDate  string  `db:"OrderDate"` // or time.Time if you prefer
	Total      float64 `db:"Total"`
}

type Product struct {
	ProductID int     `db:"ProductID"`
	Name      string  `db:"Name"`
	Price     float64 `db:"Price"`
}

type TestDataResult struct {
	Customers     []Customer
	Orders        []Order
	ProductsEmpty []Product
	Products      []Product
}

type Customer struct {
	CustomerID int    `db:"CustomerID"`
	Name       string `db:"Name"`
	// Country    string `db:"Country"`
}
