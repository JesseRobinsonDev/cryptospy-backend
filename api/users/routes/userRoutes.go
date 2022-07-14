package routes

import (
	"cryptospy-backend/api/users/controllers"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(g *gin.Engine) {
	g.POST("/user/register", controllers.RegisterUser)
	g.GET("/user/login", controllers.LoginUser)
	g.GET("/user/get/:id", controllers.GetUser)
	g.DELETE("/user/delete/:id", controllers.DeleteUser)
	g.DELETE("/user/delete", controllers.DeleteUserByName)
	g.PUT("/user/:id/:coin/track", controllers.TrackCoin)
	g.PUT("/user/:id/:coin/untrack", controllers.UntrackCoin)
}