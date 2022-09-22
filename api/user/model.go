package user

type Login struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	IsRemember bool   `json:"isRemember"`
}

type LoginRequest struct {
	Login Login `json:"login"`
}
