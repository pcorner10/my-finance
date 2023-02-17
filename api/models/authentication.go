package models

type AuthenticationInput struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"pass_word" binding:"required"`
}
