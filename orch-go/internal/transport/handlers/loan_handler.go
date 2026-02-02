package handlers

import (
	"orch-go/internal/services"
	"orch-go/internal/transport/midleware"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: В LoanService не хватает методов для получения кредита и платежей по кредиту
// с проверкой принадлежности клиенту. Пришлось делать дополнительные запросы.

func GetLoansByCustomerIdHandler(loanService services.LoanService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}

		loans, err := loanService.GetLoansByCustomer(c.Request.Context(), customerId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, loans)
	}
}

func GetLoanByIdHandler(loanService services.LoanService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}

		loanId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid loan id"})
			return
		}

		loan, err := loanService.GetLoanById(c.Request.Context(), int32(loanId))
		if err != nil {
			c.JSON(404, gin.H{"error": "loan not found"})
			return
		}

		if *loan.CustomerID != customerId {
			c.JSON(403, gin.H{"error": "you are not authorized to view this loan"})
			return
		}

		c.JSON(200, loan)
	}
}

func GetLoanPaymentsByLoanIdHandler(loanService services.LoanService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}

		loanId, err := strconv.Atoi(c.Param("loanId"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid loan id"})
			return
		}

		loan, err := loanService.GetLoanById(c.Request.Context(), int32(loanId))
		if err != nil {
			c.JSON(404, gin.H{"error": "loan not found"})
			return
		}

		if *loan.CustomerID != customerId {
			c.JSON(403, gin.H{"error": "you are not authorized to view this loan's payments"})
			return
		}

		payments, err := loanService.GetLoanPaymentsByLoan(c.Request.Context(), int32(loanId))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, payments)
	}
}

func InitLoanRouter(r *gin.Engine, loanService services.LoanService) {
	loanGroup := r.Group("/loan")
	{
		authMiddleware := midleware.NewAuthMiddleware(os.Getenv("JWT_SECRET_KEY"))
		loanGroup.Use(authMiddleware.AuthRequired())
		{
			loanGroup.GET("/", GetLoansByCustomerIdHandler(loanService))
			loanGroup.GET("/:id", GetLoanByIdHandler(loanService))
			loanGroup.GET("/:loanId/payments", GetLoanPaymentsByLoanIdHandler(loanService))
		}
	}
}
