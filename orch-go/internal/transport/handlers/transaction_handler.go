package handlers

import (
	"orch-go/internal/domain/transaction"
	"orch-go/internal/services"
	"orch-go/internal/transport/midleware"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// TODO: В TransactionService не хватает методов для получения транзакции
// с проверкой принадлежности клиенту. Пришлось делать дополнительные запросы в AccountService.

// CreateTransactionHandler godoc
// @Summary      Create transaction
// @Description  Creates a new transaction from the authenticated customer's account to a target account
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      transaction.Transaction  true  "Transaction details"
// @Success      201      {object}  transaction.Transaction
// @Failure      400      {object}  map[string]string "Bad Request"
// @Failure      403      {object}  map[string]string "Forbidden"
// @Failure      404      {object}  map[string]string "Not Found"
// @Failure      500      {object}  map[string]string "Internal Server Error"
// @Router       /transaction/ [post]
func CreateTransactionHandler(transactionService services.TransactionService, accountService services.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}

		var tx transaction.Transaction
		if err := c.ShouldBindJSON(&tx); err != nil {
			c.JSON(400, gin.H{"error": "invalid request body: " + err.Error()})
			return
		}

		// Security check: ensure the source account belongs to the authenticated user
		sourceAccount, err := accountService.GetAccountById(c.Request.Context(), *tx.SourceAccount)
		if err != nil {
			c.JSON(404, gin.H{"error": "source account not found"})
			return
		}

		if sourceAccount.CustomerId != customerId {
			c.JSON(403, gin.H{"error": "you are not authorized to perform transactions from this account"})
			return
		}
		createdAt := time.Now()
		// Set creation timestamp
		tx.CreatedAt = &createdAt

		createdTx, err := transactionService.CreateTransaction(c.Request.Context(), &tx)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to create transaction: " + err.Error()})
			return
		}

		c.JSON(201, createdTx)
	}
}

// GetTransactionsByAccountIdHandler godoc
// @Summary      Get transactions by account ID in date range
// @Description  Retrieves transactions for a given account within a specific date range. The account must belong to the authenticated customer.
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        accountId  path      int  true  "Account ID"
// @Param        request    body      transaction.GetByDateRange  true  "Date range filters"
// @Success      200        {array}   transaction.Transaction
// @Failure      400        {object}  map[string]string "Bad Request"
// @Failure      403        {object}  map[string]string "Forbidden"
// @Failure      404        {object}  map[string]string "Not Found"
// @Failure      500        {object}  map[string]string "Internal Server Error"
// @Router       /transaction/account/{accountId} [post]
func GetTransactionsByAccountIdHandler(transactionService services.TransactionService, accountService services.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}

		accountId, err := strconv.Atoi(c.Param("accountId"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid account id"})
			return
		}

		account, err := accountService.GetAccountById(c.Request.Context(), int32(accountId))
		if err != nil {
			c.JSON(404, gin.H{"error": "account not found"})
			return
		}

		if account.CustomerId != customerId {
			c.JSON(403, gin.H{"error": "you are not authorized to view these transactions"})
			return
		}

		var req transaction.GetByDateRange
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		transactions, err := transactionService.GetTransactionRevenue(c.Request.Context(), int32(accountId), req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, transactions)
	}
}

// GetTransactionByIdHandler godoc
// @Summary      Get transaction by ID
// @Description  Retrieves details of a specific transaction by ID, provided the authenticated customer owns either the source or target account
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Transaction ID"
// @Success      200  {object}  transaction.Transaction
// @Failure      400  {object}  map[string]string "Bad Request"
// @Failure      403  {object}  map[string]string "Forbidden"
// @Failure      404  {object}  map[string]string "Not Found"
// @Router       /transaction/{id} [get]
func GetTransactionByIdHandler(transactionService services.TransactionService, accountService services.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}

		transactionId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid transaction id"})
			return
		}

		tx, err := transactionService.GetTransactionById(c.Request.Context(), int32(transactionId))
		if err != nil {
			c.JSON(404, gin.H{"error": "transaction not found"})
			return
		}

		// Check if user owns the source account
		sourceAccount, err := accountService.GetAccountById(c.Request.Context(), *tx.SourceAccount)
		if err == nil && sourceAccount.CustomerId == customerId {
			c.JSON(200, tx)
			return
		}

		// Check if user owns the destination account
		destAccount, err := accountService.GetAccountById(c.Request.Context(), *tx.TargetAccount)
		if err == nil && destAccount.CustomerId == customerId {
			c.JSON(200, tx)
			return
		}

		// If neither, user is not authorized
		c.JSON(403, gin.H{"error": "you are not authorized to view this transaction"})
	}
}

// GetAllTransactionCategoriesHandler godoc
// @Summary      Get all transaction categories
// @Description  Retrieves a paginated list of all transaction categories
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        pageN     query     int   false  "Page number (default: 1)"
// @Param        pageSize  query     int   false  "Page size (default: 10)"
// @Param        orderBy   query     string  false  "Order by field (default: id)"
// @Param        isDesc    query     bool  false  "Descending order (default: false)"
// @Success      200       {array}   transaction.TransactionCategory
// @Failure      500       {object}  map[string]string "Internal Server Error"
// @Router       /transaction/categories [get]
func GetAllTransactionCategoriesHandler(s services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		pageN, _ := strconv.Atoi(c.DefaultQuery("pageN", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
		orderBy := c.DefaultQuery("orderBy", "id")
		isDesc, _ := strconv.ParseBool(c.DefaultQuery("isDesc", "false"))

		categories, err := s.GetAllTransactionCategories(c.Request.Context(), int32(pageN), int32(pageSize), orderBy, isDesc)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, categories)
	}
}

func InitTransactionRouter(r *gin.Engine, transactionService services.TransactionService, accountService services.AccountService) {
	transactionGroup := r.Group("/transaction")
	{
		authMiddleware := midleware.NewAuthMiddleware(os.Getenv("JWT_SECRET_KEY"))

		// Public route for categories
		transactionGroup.GET("/categories", GetAllTransactionCategoriesHandler(transactionService))

		transactionGroup.Use(authMiddleware.AuthRequired())
		{
			transactionGroup.POST("/", CreateTransactionHandler(transactionService, accountService))
			transactionGroup.POST("/account/:accountId", GetTransactionsByAccountIdHandler(transactionService, accountService))
			transactionGroup.GET("/:id", GetTransactionByIdHandler(transactionService, accountService))
		}
	}
}
