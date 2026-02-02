package handlers

import (
	"orch-go/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBranchByIdHandler(s services.BranchService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}

		branch, err := s.GetBranchById(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "branch not found"})
			return
		}

		c.JSON(200, branch)
	}
}

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
