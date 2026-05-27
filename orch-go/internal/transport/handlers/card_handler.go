package handlers

import (
	_ "orch-go/internal/domain/card"
	"orch-go/internal/services"
	"orch-go/internal/transport/midleware"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCardsByAccountIdHandler godoc
// @Summary      Get cards by account ID
// @Description  Retrieves all cards associated with a specific account. The account must belong to the authenticated customer.
// @Tags         cards
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        accountId  path      int  true  "Account ID"
// @Success      200        {array}   card.Card
// @Failure      400        {object}  map[string]string "Bad Request"
// @Failure      403        {object}  map[string]string "Forbidden"
// @Failure      404        {object}  map[string]string "Not Found"
// @Failure      500        {object}  map[string]string "Internal Server Error"
// @Router       /card/account/{accountId} [get]
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

		acc, err := accountService.GetAccountById(c.Request.Context(), int32(accountId))
		if err != nil {
			c.JSON(404, gin.H{"error": "account not found"})
			return
		}

		if acc.CustomerId != customerId {
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

// GetCardByIdHandler godoc
// @Summary      Get card by ID
// @Description  Retrieves details of a specific card by its ID, provided the card's account belongs to the authenticated customer.
// @Tags         cards
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Card ID"
// @Success      200  {object}  card.Card
// @Failure      400  {object}  map[string]string "Bad Request"
// @Failure      403  {object}  map[string]string "Forbidden"
// @Failure      404  {object}  map[string]string "Not Found"
// @Router       /card/{id} [get]
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

		cardVal, err := cardService.GetCardById(c.Request.Context(), int32(cardId))
		if err != nil {
			c.JSON(404, gin.H{"error": "card not found"})
			return
		}

		acc, err := accountService.GetAccountById(c.Request.Context(), *cardVal.AccountID)
		if err != nil {
			c.JSON(404, gin.H{"error": "account not found"})
			return
		}

		if acc.CustomerId != customerId {
			c.JSON(403, gin.H{"error": "you are not authorized to view this card"})
			return
		}

		c.JSON(200, cardVal)
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
