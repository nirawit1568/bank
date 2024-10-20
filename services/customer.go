package services

import "github.com/google/uuid"

type CustomerResp struct {
	CustomerID uuid.UUID `json:"customer_id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResp, error)
	GetCustomer(uuid.UUID) (*CustomerResp, error)
}
