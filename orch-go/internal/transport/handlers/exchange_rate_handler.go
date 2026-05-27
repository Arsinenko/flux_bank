package handlers

import (
	_ "orch-go/internal/domain/exchange_rate"
	"orch-go/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetExchangeRatesByBaseCurrencyHandler godoc
// @Summary      Get exchange rates by base currency
// @Description  Retrieves all exchange rates matching the specified base currency
// @Tags         exchange-rates
// @Accept       json
// @Produce      json
// @Param        baseCurrency  path      string  true  "Base Currency Code (e.g. USD)"
// @Success      200           {array}   exchange_rate.ExchangeRate
// @Failure      400           {object}  map[string]string "Bad Request"
// @Failure      500           {object}  map[string]string "Internal Server Error"
// @Router       /exchange-rate/{baseCurrency} [get]
func GetExchangeRatesByBaseCurrencyHandler(s services.ExchangeRateService) gin.HandlerFunc {
	return func(c *gin.Context) {
		baseCurrency := c.Param("baseCurrency")
		if baseCurrency == "" {
			c.JSON(400, gin.H{"error": "base currency is required"})
			return
		}

		rates, err := s.GetExchangeRatesByBaseCurrency(c.Request.Context(), baseCurrency)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, rates)
	}
}

// GetAllExchangeRatesHandler godoc
// @Summary      Get all exchange rates
// @Description  Retrieves a paginated list of all exchange rates
// @Tags         exchange-rates
// @Accept       json
// @Produce      json
// @Param        pageN     query     int   false  "Page number (default: 1)"
// @Param        pageSize  query     int   false  "Page size (default: 10)"
// @Param        orderBy   query     string  false  "Order by field (default: id)"
// @Param        isDesc    query     bool  false  "Descending order (default: false)"
// @Success      200       {array}   exchange_rate.ExchangeRate
// @Failure      500       {object}  map[string]string "Internal Server Error"
// @Router       /exchange-rate/ [get]
func GetAllExchangeRatesHandler(s services.ExchangeRateService) gin.HandlerFunc {
	return func(c *gin.Context) {
		pageN, _ := strconv.Atoi(c.DefaultQuery("pageN", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
		orderBy := c.DefaultQuery("orderBy", "id")
		isDesc, _ := strconv.ParseBool(c.DefaultQuery("isDesc", "false"))

		rates, err := s.GetAllExchangeRates(c.Request.Context(), int32(pageN), int32(pageSize), orderBy, isDesc)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, rates)
	}
}

func InitExchangeRateRouter(r *gin.Engine, exchangeRateService services.ExchangeRateService) {
	exchangeRateGroup := r.Group("/exchange-rate")
	{
		exchangeRateGroup.GET("/", GetAllExchangeRatesHandler(exchangeRateService))
		exchangeRateGroup.GET("/:baseCurrency", GetExchangeRatesByBaseCurrencyHandler(exchangeRateService))
	}
}
