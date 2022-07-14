package models

// Used for validating the body fields in "API/user/register"
type RegisterUserModel struct {
	Username string `json:"username"`
	Pass     string `json:"pass"`
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
	User_ID       int      `json:"user_id"`
	Username      string   `json:"username"`
	Tracked_Coins []string `json:"tracked_coins"`
}