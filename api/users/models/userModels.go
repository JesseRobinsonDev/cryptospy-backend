package models

type RegisterUserModel struct {
	Username string `json:"username"`
	Pass     string `json:"pass"`
}

type LoginUserRequestModel struct {
	Username string `json:"username"`
	Pass     string `json:"pass"`
}

type LoginUserResponseModel struct {
	User_ID int    `json:"user_id"`
	Message string `json:"message"`
}

type GetUserModel struct {
	User_ID       int      `json:"user_id"`
	Username      string   `json:"username"`
	Tracked_Coins []string `json:"tracked_coins"`
}