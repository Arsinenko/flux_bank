package handlers

import (
	"orch-go/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAtmByIdHandler(s services.AtmService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		atm, err := s.GetAtmById(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, atm)
	}
}

func GetAtmsByStatusHandler(s services.AtmService) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Param("status")
		atms, err := s.GetAtmsByStatus(c.Request.Context(), status)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, atms)
	}
}

func GetAtmsByLocationSubStrHandler(s services.AtmService) gin.HandlerFunc {
	return func(c *gin.Context) {
		subStr := c.Query("q")
		atms, err := s.GetAtmsByLocationSubStr(c.Request.Context(), subStr)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, atms)
	}
}

func GetAtmsByBranchHandler(s services.AtmService) gin.HandlerFunc {
	return func(c *gin.Context) {
		branchId, err := strconv.Atoi(c.Param("branchId"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid branch id"})
			return
		}
		atms, err := s.GetAtmsByBranch(c.Request.Context(), int32(branchId))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, atms)
	}
}

func GetAllAtmsHandler(s services.AtmService) gin.HandlerFunc {
	return func(c *gin.Context) {
		pageN, _ := strconv.Atoi(c.DefaultQuery("pageN", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
		orderBy := c.DefaultQuery("orderBy", "id")
		isDesc, _ := strconv.ParseBool(c.DefaultQuery("isDesc", "false"))

		atms, err := s.GetAllAtms(c.Request.Context(), int32(pageN), int32(pageSize), orderBy, isDesc)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, atms)
	}
}

func InitAtmRouter(r *gin.Engine, atmService services.AtmService) {
	atmGroup := r.Group("/atm")
	{
		atmGroup.GET("/:id", GetAtmByIdHandler(atmService))
		atmGroup.GET("/status/:status", GetAtmsByStatusHandler(atmService))
		atmGroup.GET("/location", GetAtmsByLocationSubStrHandler(atmService))
		atmGroup.GET("/branch/:branchId", GetAtmsByBranchHandler(atmService))
		atmGroup.GET("/", GetAllAtmsHandler(atmService))
	}
}
