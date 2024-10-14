package users

import "time"

type User struct {
	UserID      string `json:"user_id"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:""`
	Address     string `json:"address"`
	Pin         string `json:"pin" validate:"required,len=6"`
	Balance     int    `json:"balance"`
	CreatedAt   time.Time
}

type ResponseUser struct {
	UserID      string    `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreatedDate time.Time `json:"created_date"`
}

type ResponseRegisterSuccess struct {
	Status string       `json:"status"`
	Result ResponseUser `json:"result"`
}

type RequestLogin struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Pin         string `json:"pin" validate:"required"`
}

type ResponseLoginSuccess struct {
	Status string      `json:"status"`
	Result ResultToken `json:"result"`
}

type ResultToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
