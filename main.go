package main

import (
	"fmt"
	"log"

	_ "usermanagementservice/docs"
	"usermanagementservice/routes"

	"github.com/gin-gonic/gin"
	"github.com/honeycombio/otel-config-go/otelconfig"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title UserManagment APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @Security JWT
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
// @schemes http https
func main() {
	// @schemes http https
	// Create a new gin instance

	r := gin.Default()

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	fmt.Println(err)
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry(
		otelconfig.WithMetricsEnabled(false),
	)
	if err != nil {
		log.Fatalf("error setting up OTel SDK - %e", err)
	}

	defer otelShutdown()

	// Load the routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.PublicRoutes(r)
	routes.LoginRoutes(r)
	routes.CustomerRoutes(r)
	// Run the server
	r.Run(":8000")
}
