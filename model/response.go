package model

type MainResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type RegisterResponse struct {
	MainResponse
	UserId uint `json:"user_id,omitempty"`
}

type LoginResponse struct {
	MainResponse
	AccessToken string `json:"access_token,omitempty"`
}
