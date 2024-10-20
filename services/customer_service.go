package services

import (
	"bank/repository"
	"log"

	"github.com/google/uuid"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) *customerService {
	return &customerService{custRepo: custRepo}
}

func (s *customerService) GetCustomers() ([]CustomerResp, error) {

	customers, err := s.custRepo.GetAll()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	resp := []CustomerResp{}
	for _, v := range customers {
		resp = append(resp, CustomerResp{
			CustomerID: v.CustomerID,
			Name:       v.Name,
			Status:     v.Status,
		})
	}

	return resp, nil
}

func (s *customerService) GetCustomer(id uuid.UUID) (*CustomerResp, error) {

	customer, err := s.custRepo.GetById(id)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	resp := CustomerResp{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &resp, nil
}
