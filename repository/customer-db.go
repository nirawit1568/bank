package repository

import (
	"bank/config"
	"database/sql"

	"github.com/google/uuid"
)

type customerRepositoryDB struct {
	db *sql.DB
}

func NewCustomerRepositoryDB(db *sql.DB) *customerRepositoryDB {
	return &customerRepositoryDB{db: db}
}

func (c *customerRepositoryDB) GetAll() ([]Customer, error) {

	db := config.PgConnect()

	cusRows, err := db.Query(`select * from customers`)
	if err != nil {
		return nil, err
	}
	defer cusRows.Close()

	var customers []Customer
	for cusRows.Next() {
		var customer Customer
		if err := cusRows.Scan(&customer.CustomerID, &customer.Name, &customer.Status); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (c *customerRepositoryDB) GetById(id uuid.UUID) (*Customer, error) {

	db := config.PgConnect()

	cusRows := db.QueryRow(`select * from customers where customer_id = $1`, id)

	var customer Customer
	if err := cusRows.Scan(&customer.CustomerID, &customer.Name, &customer.Status); err != nil {
		return nil, err
	}
	return &customer, nil

}
