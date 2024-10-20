package main

import (
	"bank/config"
	"database/sql"
	"fmt"
	"log"
)

func main() {

	db := config.PgConnect()
	log.Println("Starting database migration...")

	// ตัวอย่างการสร้างตารางใหม่
	if err := migrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully!")
}

func migrate(db *sql.DB) error {
	query := `
		-- Create the customers table
		CREATE TABLE customers (
			customer_id UUID PRIMARY KEY,  -- UUID for customer_id as PK
			name VARCHAR(255) NOT NULL,    -- Name of the customer (string)
			status INT                     -- Status of the customer (int64)
		);

		-- Create the accounts table
		CREATE TABLE accounts (
			id UUID PRIMARY KEY,           -- UUID for account ID as PK
			account_id VARCHAR(255) NOT NULL, -- Unique account ID (string)
			customer_id UUID,              -- Foreign key linking to customer_id in customers table
			opening_date TIMESTAMPTZ,      -- Timestamp with time zone for the opening date
			account_type VARCHAR(50),      -- Type of account (string)
			amount NUMERIC(15, 2),         -- Amount with two decimal precision (float64)
			status INT,                    -- Status of the account (int64)
			
			-- Set up foreign key constraint linking to customers table
			CONSTRAINT fk_customer
			FOREIGN KEY(customer_id) 
			REFERENCES customers(customer_id)
		);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Accounts and Customers was created.")
	return nil
}
