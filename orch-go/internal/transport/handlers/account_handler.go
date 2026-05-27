package handlers

import (
	"github.com/shopspring/decimal"
	"orch-go/internal/domain/account"
	"orch-go/internal/services"
	"orch-go/internal/transport/midleware"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateAccountHandler godoc
// @Summary      Create account
// @Description  Creates a new account with default parameters (type_id=1, balance=100) for the authenticated customer
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  account.Account
// @Failure      400  {object}  map[string]string "Bad Request"
// @Failure      401  {object}  map[string]string "Unauthorized"
// @Router       /account [post]
func CreateAccountHandler(s services.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}
		acc, err := s.CreateAccount(c.Request.Context(), &account.Account{
			Id:         nil,
			CustomerId: customerId,
			TypeId:     1,
			Iban:       "Account",
			Balance:    decimal.NewFromInt(100),
			CreatedAt:  time.Now(),
			IsActive:   true,
		})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, acc)
		return
	}
}

// GetUserAccountsHandler godoc
// @Summary      Get user accounts
// @Description  Retrieves all accounts belonging to the authenticated customer
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   account.Account
// @Failure      400  {object}  map[string]string "Bad Request"
// @Failure      401  {object}  map[string]string "Unauthorized"
// @Router       /account [get]
func GetUserAccountsHandler(s services.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}

		accounts, err := s.GetAccountsByCustomer(c.Request.Context(), customerId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, accounts)
		return
	}
}

// GetAccountsIdsByCustomer godoc
// @Summary      Get account IDs by customer ID
// @Description  Retrieves a list of account IDs associated with a specific customer ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Customer ID"
// @Success      200  {array}   int32
// @Failure      400  {object}  map[string]string "Bad Request"
// @Router       /account/customer/{id}/ids [get]
func GetAccountsIdsByCustomer(s services.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		accounts, err := s.GetAccountsByCustomer(c.Request.Context(), int32(customerId))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ids := make([]int32, len(accounts))
		for i, acc := range accounts {
			ids[i] = *acc.Id
		}
		c.JSON(200, ids)
		return
	}
}

// UpdateAccountHandler godoc
// @Summary      Update account
// @Description  Updates an existing account. The customer ID of the request must match the customer ID of the database account.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        account  body      account.Account  true  "Account data to update"
// @Success      200      {object}  map[string]string "account updated"
// @Failure      400      {object}  map[string]string "Bad Request"
// @Router       /account [put]
func UpdateAccountHandler(s services.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var model account.Account
		if err := c.ShouldBindJSON(&model); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		dbAccount, err := s.GetAccountById(c.Request.Context(), *model.Id)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if dbAccount.CustomerId != model.CustomerId {
			c.JSON(400, gin.H{"error": "customer id does not match. You can get ban for this action."})
			return
		}

		err = s.UpdateAccount(c.Request.Context(), &model)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "account updated"})
		return
	}
}

// GetAccountByIdHandler godoc
// @Summary      Get account by ID
// @Description  Retrieves an account by its ID, but only if it belongs to the authenticated customer
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  account.Account
// @Failure      400  {object}  map[string]string "Bad Request"
// @Failure      401  {object}  map[string]string "Unauthorized"
// @Router       /account/{id} [get]
func GetAccountByIdHandler(s services.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		acc, err := s.GetAccountById(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if acc.CustomerId != customerId {
			c.JSON(400, "not your account")
			return
		}
		c.JSON(200, acc)
		return
	}
}

func InitAccountRouter(r *gin.Engine, accountService services.AccountService) {
	accountGroup := r.Group("/account")
	{
		authMiddleware := midleware.NewAuthMiddleware(os.Getenv("JWT_SECRET_KEY"))
		accountGroup.Use(authMiddleware.AuthRequired())
		{
			accountGroup.POST("/", CreateAccountHandler(accountService))
			accountGroup.GET("/", GetUserAccountsHandler(accountService))
			accountGroup.GET("/:id", GetAccountByIdHandler(accountService))
			accountGroup.PUT("/", UpdateAccountHandler(accountService))
		}
	}
}
