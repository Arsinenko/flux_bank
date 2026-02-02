package handlers

import (
	"orch-go/internal/services"
	"orch-go/internal/transport/midleware"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCardsByAccountIdHandler(cardService services.CardService, accountService services.AccountService) gin.HandlerFunc {
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
			c.JSON(403, gin.H{"error": "you are not authorized to view these cards"})
			return
		}

		cards, err := cardService.GetCardsByAccountId(c.Request.Context(), int32(accountId))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, cards)
	}
}

func GetCardByIdHandler(cardService services.CardService, accountService services.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}

		cardId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid card id"})
			return
		}

		card, err := cardService.GetCardById(c.Request.Context(), int32(cardId))
		if err != nil {
			c.JSON(404, gin.H{"error": "card not found"})
			return
		}

		account, err := accountService.GetAccountById(c.Request.Context(), *card.AccountID)
		if err != nil {
			c.JSON(404, gin.H{"error": "account not found"})
			return
		}

		if account.CustomerId != customerId {
			c.JSON(403, gin.H{"error": "you are not authorized to view this card"})
			return
		}

		c.JSON(200, card)
	}
}

func InitCardRouter(r *gin.Engine, cardService services.CardService, accountService services.AccountService) {
	cardGroup := r.Group("/card")
	{
		authMiddleware := midleware.NewAuthMiddleware(os.Getenv("JWT_SECRET_KEY"))
		cardGroup.Use(authMiddleware.AuthRequired())
		{
			cardGroup.GET("/account/:accountId", GetCardsByAccountIdHandler(cardService, accountService))
			cardGroup.GET("/:id", GetCardByIdHandler(cardService, accountService))
		}
	}
}
