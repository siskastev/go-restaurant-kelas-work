package model

type User struct {
	ID       string `json:"id"`
	Username string `gorm:"column:username;unique" json:"username"`
	Hash     string `gorm:"column:hash" json:"-"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSession struct {
	JWTToken string `json:"jwt_token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
