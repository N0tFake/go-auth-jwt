package routes

import (
	"authJWT/controllers"
	"authJWT/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/protected")
	protected.Use(middlewares.AuthMiddleware())
	protected.GET("/welcome", Welcome)
	protected.GET("/users", controllers.Users)

	return r
}

func Welcome(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{"data": "Welcome " + username})
}
