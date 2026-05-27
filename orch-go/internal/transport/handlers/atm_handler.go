package handlers

import (
	_ "orch-go/internal/domain/atm"
	"orch-go/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAtmByIdHandler godoc
// @Summary      Get ATM by ID
// @Description  Retrieves details of a specific ATM by its ID
// @Tags         atms
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ATM ID"
// @Success      200  {object}  atm.Atm
// @Failure      400  {object}  map[string]string "Bad Request"
// @Router       /atm/{id} [get]
func GetAtmByIdHandler(s services.AtmService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		atmVal, err := s.GetAtmById(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, atmVal)
	}
}

// GetAtmsByStatusHandler godoc
// @Summary      Get ATMs by status
// @Description  Retrieves a list of ATMs filtered by status
// @Tags         atms
// @Accept       json
// @Produce      json
// @Param        status   path      string  true  "ATM Status"
// @Success      200      {array}   atm.Atm
// @Failure      400      {object}  map[string]string "Bad Request"
// @Router       /atm/status/{status} [get]
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

// GetAtmsByLocationSubStrHandler godoc
// @Summary      Get ATMs by location substring
// @Description  Retrieves a list of ATMs where the location matches a substring query
// @Tags         atms
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "Location query string"
// @Success      200  {array}   atm.Atm
// @Failure      400  {object}  map[string]string "Bad Request"
// @Router       /atm/location [get]
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

// GetAtmsByBranchHandler godoc
// @Summary      Get ATMs by branch ID
// @Description  Retrieves all ATMs located at a specific branch
// @Tags         atms
// @Accept       json
// @Produce      json
// @Param        branchId  path      int  true  "Branch ID"
// @Success      200       {array}   atm.Atm
// @Failure      400       {object}  map[string]string "Bad Request"
// @Router       /atm/branch/{branchId} [get]
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

// GetAllAtmsHandler godoc
// @Summary      Get all ATMs
// @Description  Retrieves a paginated list of all ATMs
// @Tags         atms
// @Accept       json
// @Produce      json
// @Param        pageN     query     int   false  "Page number (default: 1)"
// @Param        pageSize  query     int   false  "Page size (default: 10)"
// @Param        orderBy   query     string  false  "Order by field (default: id)"
// @Param        isDesc    query     bool  false  "Descending order (default: false)"
// @Success      200       {array}   atm.Atm
// @Failure      400       {object}  map[string]string "Bad Request"
// @Router       /atm/ [get]
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
