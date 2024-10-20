package repository

import (
	"github.com/google/uuid"
)

type Customer struct {
	CustomerID uuid.UUID `db:"customer_id"`
	Name       string    `db:"name"`
	Status     string    `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(uuid.UUID) (*Customer, error)
}
