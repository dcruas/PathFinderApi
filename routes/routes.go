package routes

import (
	"pathfinderapi/handlers/skillcheckhdl"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/pathfinder2e/v1/distribution", skillcheckhdl.GetCheckStatics)
	}
}
