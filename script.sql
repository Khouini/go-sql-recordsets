-- 1. Create a new database
CREATE DATABASE TestDB_X;
GO

-- 2. Use the new database
USE TestDB_X;
GO

-- 3. Create Customers table
CREATE TABLE Customers (
                           CustomerID INT PRIMARY KEY,
                           Name VARCHAR(100),
                           Country VARCHAR(50)
);
GO

-- 4. Create Orders table
CREATE TABLE Orders (
                        OrderID INT PRIMARY KEY,
                        CustomerID INT,
                        OrderDate DATE,
                        Total DECIMAL(10, 2),
                        FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID)
);
GO

-- 5. Create Products table
CREATE TABLE Products (
                          ProductID INT PRIMARY KEY,
                          Name VARCHAR(100),
                          Price DECIMAL(10, 2)
);
GO

-- 6. Insert test data into Customers
INSERT INTO Customers (CustomerID, Name, Country) VALUES
                                                      (1, 'Alice', 'USA'),
                                                      (2, 'Bob', 'UK'),
                                                      (3, 'Carlos', 'Spain');
GO

-- 7. Insert test data into Orders
INSERT INTO Orders (OrderID, CustomerID, OrderDate, Total) VALUES
                                                               (101, 1, '2025-07-01', 250.00),
                                                               (102, 2, '2025-07-02', 125.50),
                                                               (103, 1, '2025-07-02', 99.99);
GO

-- 8. Insert test data into Products
INSERT INTO Products (ProductID, Name, Price) VALUES
                                                  (1001, 'Laptop', 1000.00),
                                                  (1002, 'Mouse', 25.00),
                                                  (1003, 'Keyboard', 45.00);
GO

-- 9. Create a stored procedure to return 3 recordsets
CREATE PROCEDURE sp_GetTestData
AS
BEGIN
    SELECT * FROM Customers;
    SELECT * FROM Orders;
    SELECT * FROM Products;
END;
GO

CREATE PROCEDURE sp_GetCustomers
AS
BEGIN
    SELECT * FROM Customers;
END;
GO

-- 10. Execute the stored procedure (optional)
EXEC sp_GetTestData;
