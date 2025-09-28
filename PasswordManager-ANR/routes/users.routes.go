package routes

import (
	"github.com/anarrod172/passwordmanager/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/users", controllers.GetUser)
	router.POST("/users/newuser", controllers.PostUser)
	router.DELETE("/users/username/:username", controllers.DeleteUser)
	router.PUT("/users/id/:id", controllers.PutNotes)

	//router.Run("localhost:8080")
	return router
}
