package models

import "time"

// Used for validating the body fields in "API/user/register"
type RegisterUserRequestModel struct {
	Username string `json:"username"`
	Pass     string `json:"pass"`
}

// Used on successful response in "API/user/register"
type RegisterUserResponseModel struct {
	User_ID int    `json:"user_id"`
	Message string `json:"message"`
}

// Used for validating the body fields in "API/user/login"
type LoginUserRequestModel struct {
	Username string `json:"username"`
	Pass     string `json:"pass"`
}

// Used on successful response in "API/user/login"
type LoginUserResponseModel struct {
	User_ID int    `json:"user_id"`
	Message string `json:"message"`
}

// Used on successful response in "API/user/:id/get"
type GetUserModel struct {
	User_ID       int       `json:"user_id"`
	Username      string    `json:"username"`
	Tracked_Coins []string  `json:"tracked_coins"`
	Date_Joined   time.Time `json:"date_joined"`
}