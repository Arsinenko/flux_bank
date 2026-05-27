package handlers

import (
	_ "orch-go/internal/domain/deposit"
	"orch-go/internal/services"
	"orch-go/internal/transport/midleware"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: В DepositService не хватает метода для получения депозита по ID с проверкой принадлежности клиенту.
// Пришлось делать дополнительный запрос в сервис для проверки.

// GetDepositsByCustomerIdHandler godoc
// @Summary      Get customer deposits
// @Description  Retrieves all deposits associated with the authenticated customer
// @Tags         deposits
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   deposit.Deposit
// @Failure      401  {object}  map[string]string "Unauthorized"
// @Failure      500  {object}  map[string]string "Internal Server Error"
// @Router       /deposit/ [get]
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

// GetDepositByIdHandler godoc
// @Summary      Get deposit by ID
// @Description  Retrieves details of a specific deposit by its ID, provided it belongs to the authenticated customer
// @Tags         deposits
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Deposit ID"
// @Success      200  {object}  deposit.Deposit
// @Failure      400  {object}  map[string]string "Bad Request"
// @Failure      403  {object}  map[string]string "Forbidden"
// @Failure      404  {object}  map[string]string "Not Found"
// @Router       /deposit/{id} [get]
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

		depositVal, err := depositService.GetDepositById(c.Request.Context(), int32(depositId))
		if err != nil {
			c.JSON(404, gin.H{"error": "deposit not found"})
			return
		}

		if depositVal.CustomerID != customerId {
			c.JSON(403, gin.H{"error": "you are not authorized to view this deposit"})
			return
		}

		c.JSON(200, depositVal)
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
