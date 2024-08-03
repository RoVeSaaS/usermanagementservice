package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health CHeck godoc
// @Summary Health Check Endpoint for the service.
// @Description Health Check Endpoint for the service.
// @Tags HealthCheck
// @Produce json
// @Success 200
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Success": "Health Check Passed"})
}
