package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"orch-go/internal/services"
	"orch-go/internal/simulation/simulation_runner"
)

// StartSimulationHandler godoc
// @Summary      Start simulation
// @Description  Starts or resumes the simulation loop with a specified or default agent count
// @Tags         simulation
// @Accept       json
// @Produce      json
// @Param        count  query     int  false  "Agent Count"
// @Success      200    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]string "Internal Server Error"
// @Router       /simulation/start [post]
func StartSimulationHandler(container *services.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var count int
		countStr := c.Query("count")
		if countStr != "" {
			var err error
			count, err = strconv.Atoi(countStr)
			if err != nil || count <= 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid count parameter"})
				return
			}
		}

		err := simulationrunner.DefaultManager.Start(container, count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		running, currentCount := simulationrunner.DefaultManager.GetStatus()
		c.JSON(http.StatusOK, gin.H{
			"message": "simulation started",
			"running": running,
			"count":   currentCount,
		})
	}
}

// PauseSimulationHandler godoc
// @Summary      Pause simulation
// @Description  Pauses the active simulation and saves agent state to agents.json
// @Tags         simulation
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string "Internal Server Error"
// @Router       /simulation/pause [post]
func PauseSimulationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := simulationrunner.DefaultManager.Pause()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "simulation paused"})
	}
}

// SetAgentCountHandler godoc
// @Summary      Set agent count
// @Description  Sets the number of active agents. If running, restarts the simulation with the new count.
// @Tags         simulation
// @Accept       json
// @Produce      json
// @Param        count  query     int  true  "Agent Count"
// @Success      200    {object}  map[string]interface{}
// @Failure      400    {object}  map[string]string "Bad Request"
// @Failure      500    {object}  map[string]string "Internal Server Error"
// @Router       /simulation/count [post]
func SetAgentCountHandler(container *services.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		countStr := c.Query("count")
		if countStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing count parameter"})
			return
		}

		count, err := strconv.Atoi(countStr)
		if err != nil || count <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid count parameter"})
			return
		}

		err = simulationrunner.DefaultManager.SetAgentCount(container, count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		running, currentCount := simulationrunner.DefaultManager.GetStatus()
		c.JSON(http.StatusOK, gin.H{
			"message": "agent count updated",
			"running": running,
			"count":   currentCount,
		})
	}
}

// GetSimulationStatusHandler godoc
// @Summary      Get simulation status
// @Description  Retrieves the running state and current agent count of the simulation
// @Tags         simulation
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /simulation/status [get]
func GetSimulationStatusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		running, count := simulationrunner.DefaultManager.GetStatus()
		c.JSON(http.StatusOK, gin.H{
			"running": running,
			"count":   count,
		})
	}
}

// InitSimulationRouter registers all simulation endpoints.
func InitSimulationRouter(r *gin.Engine, container *services.ServiceContainer) {
	simGroup := r.Group("/simulation")
	{
		simGroup.POST("/start", StartSimulationHandler(container))
		simGroup.POST("/pause", PauseSimulationHandler())
		simGroup.POST("/count", SetAgentCountHandler(container))
		simGroup.GET("/status", GetSimulationStatusHandler())
	}
}
