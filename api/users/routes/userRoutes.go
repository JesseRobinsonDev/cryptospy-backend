package routes

import (
	"cryptospy-backend/api/users/controllers"

	"github.com/gin-gonic/gin"
)

// Initializes user related routes
func InitUserRoutes(g *gin.Engine) {
	
	// LOGIN/REGISTER:
	// Takes in a Username and Password as the body parameters
	g.POST("/user/register", controllers.RegisterUser)
	// Takes in the Username and Password as the body parameters
	g.POST("/user/login", controllers.LoginUser)

	// GET:
	// Takes in the User Database ID as a url parameter
	g.GET("/user/get/:id", controllers.GetUser)

	// DELETE ACCOUNT:
	// Takes in the User Database ID as a url parameter
	g.DELETE("/user/delete/:id", controllers.DeleteUser)
	// Takes in the Username and Password as the body parameters
	g.DELETE("/user/delete", controllers.DeleteUserByName)

	// COIN TRACKING:
	// Takes in the User Database ID and Coingecko Coin ID as the url parameters
	g.PUT("/user/:id/:coin/track", controllers.TrackCoin)
	// Takes in the User Database ID and Coingecko Coin ID as the url parameters
	g.PUT("/user/:id/:coin/untrack", controllers.UntrackCoin)
}