package handlers

import (
	"orch-go/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
