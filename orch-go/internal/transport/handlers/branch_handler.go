package handlers

import (
	_ "orch-go/internal/domain/branch"
	"orch-go/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBranchByIdHandler godoc
// @Summary      Get branch by ID
// @Description  Retrieves details of a specific branch by its ID
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Branch ID"
// @Success      200  {object}  branch.Branch
// @Failure      404  {object}  map[string]string "Branch Not Found"
// @Router       /branch/{id} [get]
func GetBranchByIdHandler(s services.BranchService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}

		branchVal, err := s.GetBranchById(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "branch not found"})
			return
		}

		c.JSON(200, branchVal)
	}
}

// GetAllBranchesHandler godoc
// @Summary      Get all branches
// @Description  Retrieves a paginated list of all branches
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        page_n     query     int   false  "Page number (default: 0)"
// @Param        page_size  query     int   false  "Page size (default: 0)"
// @Param        orderBy    query     string  false  "Order by field (default: id)"
// @Param        is_desc    query     bool  false  "Descending order (default: false)"
// @Success      200        {array}   branch.Branch
// @Failure      500        {object}  map[string]string "Internal Server Error"
// @Router       /branch/ [get]
func GetAllBranchesHandler(s services.BranchService) gin.HandlerFunc {
	return func(c *gin.Context) {
		pageN, _ := strconv.Atoi(c.DefaultQuery("page_n", "0"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "0"))
		orderBy := c.DefaultQuery("orderBy", "id")
		isDesc, _ := strconv.ParseBool(c.DefaultQuery("is_desc", "false"))

		branches, err := s.GetAllBranches(c.Request.Context(), int32(pageN), int32(pageSize), orderBy, isDesc)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, branches)
	}
}

func InitBranchRouter(r *gin.Engine, branchService services.BranchService) {
	branchGroup := r.Group("/branch")
	{
		branchGroup.GET("/:id", GetBranchByIdHandler(branchService))
		branchGroup.GET("/", GetAllBranchesHandler(branchService))
	}
}
