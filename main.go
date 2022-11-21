package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"latihan1/controllers"
	"latihan1/docs"
)

// @title           Hacktiv8 Swagger
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:4444
// @BasePath  /api/v1
func main() {
	//driver={SQL Server};server=%s;user id=%s;password=%s;database=%s;encrypt=disable
	dsn := `sqlserver://sa:Pass123*@localhost/SQLSERVER14?database=GolangTraining`
	//dsn := `server=localhost\SQLSERVER14;user id=sa;password=Pass123*;database=GolangTraining;encrypt=disable`
	_, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:4444"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.New()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", controllers.Ping)
	}

	todos := v1.Group("/todos")
	{
		todos.POST("", controllers.CreateTodo)
		todos.PUT("/:id", controllers.UpdateByID)
		todos.GET("", controllers.GetAll)
		todos.GET("/:id", controllers.GetByID)
		todos.DELETE("/:id", controllers.DeleteByID)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":4444")
}
