package init

import (
	"go-sea-crm/internal/controllers"
	"go-sea-crm/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.AuthMiddleware())

	v1 := r.Group("/v1/2024/user")
	{
		v1.GET("/", controllers.NewUserController().GetUserInfo)
		// v1.POST("/", Pong)
		// v1.PUT("/", Pong)
		// v1.PATCH("/", Pong)
		// v1.DELETE("/", Pong)
	}

	// v2 := r.Group("/v2/2024/ping")
	// {
	// 	v2.GET("/", Pong)
	// 	v2.POST("/", Pong)
	// 	v2.PUT("/", Pong)
	// 	v2.PATCH("/", Pong)
	// 	v2.DELETE("/", Pong)
	// }

	return r
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping pong ok",
	})
}
