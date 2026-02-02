package handlers

import (
	"orch-go/internal/services"
	"orch-go/internal/transport/midleware"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: В DepositService не хватает метода для получения депозита по ID с проверкой принадлежности клиенту.
// Пришлось делать дополнительный запрос в сервис для проверки.

func GetDepositsByCustomerIdHandler(depositService services.DepositService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}

		deposits, err := depositService.GetDepositsByCustomer(c.Request.Context(), customerId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, deposits)
	}
}

func GetDepositByIdHandler(depositService services.DepositService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := GetIdFromRequest(c)
		if done {
			return
		}

		depositId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid deposit id"})
			return
		}

		deposit, err := depositService.GetDepositById(c.Request.Context(), int32(depositId))
		if err != nil {
			c.JSON(404, gin.H{"error": "deposit not found"})
			return
		}

		if deposit.CustomerID != customerId {
			c.JSON(403, gin.H{"error": "you are not authorized to view this deposit"})
			return
		}

		c.JSON(200, deposit)
	}
}

func InitDepositRouter(r *gin.Engine, depositService services.DepositService) {
	depositGroup := r.Group("/deposit")
	{
		authMiddleware := midleware.NewAuthMiddleware(os.Getenv("JWT_SECRET_KEY"))
		depositGroup.Use(authMiddleware.AuthRequired())
		{
			depositGroup.GET("/", GetDepositsByCustomerIdHandler(depositService))
			depositGroup.GET("/:id", GetDepositByIdHandler(depositService))
		}
	}
}
