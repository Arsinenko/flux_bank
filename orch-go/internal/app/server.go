package app

import (
	"orch-go/internal/services"
	"orch-go/internal/transport/handlers"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the Gin router and registers all REST API handlers.
func InitRouter(container *services.ServiceContainer) *gin.Engine {
	r := gin.Default()

	// Initialize all domain-specific routers
	handlers.InitAccountRouter(r, *container.AccountService)
	handlers.InitAtmRouter(r, *container.AtmService)
	handlers.InitBranchRouter(r, *container.BranchService)
	handlers.InitCardRouter(r, *container.CardService, *container.AccountService)
	handlers.InitCustomerRouter(r, *container.CustomerService, *container.UserCredentialService, *container.CustomerAddressService)
	handlers.InitDepositRouter(r, *container.DepositService)
	handlers.InitExchangeRateRouter(r, *container.ExchangeRateService)
	handlers.InitLoanRouter(r, *container.LoanService)
	handlers.InitTransactionRouter(r, *container.TransactionService, *container.AccountService)

	// Initialize Swagger documentation router
	handlers.InitSwaggerRouter(r)

	// Initialize Simulation API router
	handlers.InitSimulationRouter(r, container)

	return r
}
