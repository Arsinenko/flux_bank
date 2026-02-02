package handlers

import (
	account2 "orch-go/internal/domain/account"
	"orch-go/internal/services"
	"orch-go/internal/transport/midleware"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAccountHandler(s services.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}
		account, err := s.CreateAccount(c.Request.Context(), &account2.Account{
			Id:         nil,
			CustomerId: customerId,
			TypeId:     1,
			Iban:       "Account",
			Balance:    "100",
			CreatedAt:  time.Now(),
			IsActive:   true,
		})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, account)
		return
	}
}

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

func UpdateAccountHandler(s services.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var model account2.Account
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
		account, err := s.GetAccountById(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if account.CustomerId != customerId {
			c.JSON(400, "not your account")
			return
		}
		c.JSON(200, account)
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
