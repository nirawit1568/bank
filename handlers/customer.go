package handlers

import (
	"bank/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type customerHandler struct {
	custSrv services.CustomerService
}

func NewCustomerHandler(custSrv services.CustomerService) *customerHandler {
	return &customerHandler{custSrv: custSrv}
}

func (s *customerHandler) GetCustomers(c *fiber.Ctx) error {

	resp, err := s.custSrv.GetCustomers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (s *customerHandler) GetCustomer(c *fiber.Ctx) error {

	id := c.Params("id")
	uuidVal, err := uuid.Parse(id)
	if err != nil {
		log.Fatalf("Error parsing UUID: %v", err)
	}

	resp, err := s.custSrv.GetCustomer(uuidVal)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
