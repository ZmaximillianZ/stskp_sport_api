package main

import (
	_ "github.com/ZmaximillianZ/stskp_sport_api/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/db"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/routes"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/setting"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"log"
)

// init is invoked before main()
// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey JWT
// @in header
// @name X-AUTH-TOKEN
// @host localhost:8081
func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	settings := setting.LoadSetting()

	logging.Setup(&settings.App)
	defer func() {
		err := logging.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	dbContext, err := db.CreateDatabaseContext(settings.DBConfig)
	if err != nil {
		log.Print(err)
		panic(err)
	}

	userRepo := repository.NewUserRepository(dbContext.Connection, dbContext.QueryBuilder)
	userController := controllers.NewUserController(userRepo, validation.Validation{}, e.ErrorHandler{})

	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json") // The url pointing to API definition
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	routes.RegisterAPIV1(app.Group("/api/v1"), userController)

	if err := app.Run(":" + settings.ServerConfig.Port); err != nil {
		logging.Error(err)
	}
}
