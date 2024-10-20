package route

import (
	"bank/config"
	"bank/handlers"
	"bank/repository"
	"bank/services"

	"github.com/gofiber/fiber/v2"
)

// UserRoutes defines routes related to users
func Routes(app *fiber.App) {

	db := config.PgConnect()
	custRepo := repository.NewCustomerRepositoryDB(db)
	custServices := services.NewCustomerService(custRepo)
	custHandler := handlers.NewCustomerHandler(custServices)

	app.Get("/customers", custHandler.GetCustomers)
	app.Get("/customer/:id", custHandler.GetCustomer)
}
