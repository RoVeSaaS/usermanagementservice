package routes

import (
	"os"
	"usermanagementservice/controllers"
	middleware "usermanagementservice/middlewares"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func CustomerRoutes(r *gin.Engine) {
	r.Use(otelgin.Middleware(os.Getenv("OTEL_SERVICE_NAME")))
	r.Use(middleware.AuthenticationMiddleware())
	r.GET("/listmembers", controllers.ListMembership)
	r.POST("/inviteuser", controllers.InviteUser)
}

func PublicRoutes(r *gin.Engine) {
	r.GET("/health", controllers.HealthCheck)
}
