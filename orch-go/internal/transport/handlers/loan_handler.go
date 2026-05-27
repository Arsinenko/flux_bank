package handlers

import (
	_ "orch-go/internal/domain/loan"
	"orch-go/internal/services"
	"orch-go/internal/transport/midleware"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: В LoanService не хватает методов для получения кредита и платежей по кредиту
// с проверкой принадлежности клиенту. Пришлось делать дополнительные запросы.

// GetLoansByCustomerIdHandler godoc
// @Summary      Get customer loans
// @Description  Retrieves all loans belonging to the authenticated customer
// @Tags         loans
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   loan.Loan
// @Failure      401  {object}  map[string]string "Unauthorized"
// @Failure      500  {object}  map[string]string "Internal Server Error"
// @Router       /loan/ [get]
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

// GetLoanByIdHandler godoc
// @Summary      Get loan by ID
// @Description  Retrieves details of a specific loan by its ID, provided it belongs to the authenticated customer
// @Tags         loans
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Loan ID"
// @Success      200  {object}  loan.Loan
// @Failure      400  {object}  map[string]string "Bad Request"
// @Failure      403  {object}  map[string]string "Forbidden"
// @Failure      404  {object}  map[string]string "Not Found"
// @Router       /loan/{id} [get]
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

		loanVal, err := loanService.GetLoanById(c.Request.Context(), int32(loanId))
		if err != nil {
			c.JSON(404, gin.H{"error": "loan not found"})
			return
		}

		if *loanVal.CustomerID != customerId {
			c.JSON(403, gin.H{"error": "you are not authorized to view this loan"})
			return
		}

		c.JSON(200, loanVal)
	}
}

// GetLoanPaymentsByLoanIdHandler godoc
// @Summary      Get loan payments by loan ID
// @Description  Retrieves a list of payments made or due for a specific loan. The loan must belong to the authenticated customer.
// @Tags         loans
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        loanId  path      int  true  "Loan ID"
// @Success      200     {array}   loan.LoanPayment
// @Failure      400     {object}  map[string]string "Bad Request"
// @Failure      403     {object}  map[string]string "Forbidden"
// @Failure      404     {object}  map[string]string "Not Found"
// @Failure      500     {object}  map[string]string "Internal Server Error"
// @Router       /loan/{loanId}/payments [get]
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

		loanVal, err := loanService.GetLoanById(c.Request.Context(), int32(loanId))
		if err != nil {
			c.JSON(404, gin.H{"error": "loan not found"})
			return
		}

		if *loanVal.CustomerID != customerId {
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
