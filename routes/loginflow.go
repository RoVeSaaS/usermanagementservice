package routes

import (
	"os"
	"usermanagementservice/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	//middleware "usermanagementservice/middlewares"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func LoginRoutes(r *gin.Engine) {
	var store = cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET")))
	r.Use(sessions.Sessions("authsession", store))
	r.Use(otelgin.Middleware(os.Getenv("OTEL_SERVICE_NAME")))
	r.GET("/login", controllers.GetAuthUrl)
	r.GET("/callback", controllers.CallBack)
	r.GET("/userinfo", controllers.UserInfo)
	//r.GET("/getuserinfo", controllers.GetUserInfo)
}
