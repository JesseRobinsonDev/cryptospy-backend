package controllers

import (
	"context"
	"crypto/sha256"
	"net/http"
	"strconv"
	"strings"
	"time"

	"cryptospy-backend/api/users/models"
	"cryptospy-backend/config"

	"github.com/gin-gonic/gin"
)

// 99% of this file needs to be refactored :sad:


func convert(b []byte) string {
    s := make([]string,len(b))
    for i := range b {
        s[i] = strconv.Itoa(int(b[i]))
    }
    return strings.Join(s,"")
}

func CreateTable() {
	Conn := config.DatabaseConnect()

	sqlStr := "CREATE TABLE users (user_id SERIAL PRIMARY KEY, tracked_coins TEXT[], date_joined TIMESTAMP DEFAULT now(), username VARCHAR(16) UNIQUE, pass VARCHAR(256));"

	Conn.Exec(context.Background(), sqlStr)

	Conn.Close(context.Background())
}

func RegisterUser(c *gin.Context) {

	Conn := config.DatabaseConnect()

	var user models.RegisterUserRequestModel

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	pass := sha256.Sum256([]byte(user.Pass))
	user.Pass = convert(pass[:])

	row := Conn.QueryRow(context.Background(),
	"INSERT INTO users (username, pass, tracked_coins) VALUES ($1, $2, '{}') RETURNING user_id",
	user.Username, user.Pass)
	
	var user_id int
	if err := row.Scan(&user_id); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	Conn.Close(context.Background())

	c.JSON(http.StatusOK, models.LoginUserResponseModel{User_ID: user_id, Message: "Successfully Registered"})
}

func LoginUser(c *gin.Context) {

	var user models.LoginUserRequestModel

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	pass := sha256.Sum256([]byte(user.Pass))
	user.Pass = convert(pass[:])

	row := config.Conn.QueryRow(context.Background(),
	"SELECT user_id, pass FROM users WHERE username=$1",
	user.Username)

	var password string
	var user_id int

	if err := row.Scan(&user_id, &password); err != nil {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}

	if password != user.Pass {
		c.JSON(http.StatusUnauthorized, "Incorrect Password")
		return
	}

	c.JSON(http.StatusOK, models.LoginUserResponseModel{User_ID: user_id, Message: "Successfully Logged In"})
}

func GetUser(c *gin.Context) {

	row := config.Conn.QueryRow(context.Background(),
	"SELECT user_id, username, tracked_coins, date_joined FROM users WHERE user_id=$1",
	c.Params.ByName("id"))

	var username string
	var tracked_coins []string
	var user_id int
	var date_joined time.Time

	if err := row.Scan(&user_id, &username, &tracked_coins, &date_joined); err != nil {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}

	user := models.GetUserModel{Tracked_Coins: tracked_coins, Username: username, User_ID: user_id, Date_Joined: date_joined}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {}

func DeleteUserByName(c *gin.Context) {

	Conn := config.DatabaseConnect()

	var user models.LoginUserRequestModel

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	pass := sha256.Sum256([]byte(user.Pass))
	user.Pass = convert(pass[:])

	if _, err := Conn.Exec(context.Background(),
	"DELETE FROM users WHERE username=$1 AND pass=$2",
	user.Username, user.Pass); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "Successfully Deleted")
}

func TrackCoin(c *gin.Context) {

	Conn := config.DatabaseConnect()

	if _, err := Conn.Exec(context.Background(),
	"UPDATE users SET tracked_coins=array_append(tracked_coins, $1) WHERE user_id=$2",
	c.Params.ByName("coin"), c.Params.ByName("id")); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	Conn.Close(context.Background())

	c.JSON(http.StatusOK, "Added tracked coin")
}

func UntrackCoin(c *gin.Context) {

	Conn := config.DatabaseConnect()

	if _, err := Conn.Exec(context.Background(),
	"UPDATE users SET tracked_coins=array_remove(tracked_coins, $1) WHERE user_id=$2",
	c.Params.ByName("coin"), c.Params.ByName("id")); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	Conn.Close(context.Background())

	c.JSON(http.StatusOK, "Removed tracked coin")
}