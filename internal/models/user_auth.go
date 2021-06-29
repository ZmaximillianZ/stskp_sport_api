package models

type Auth struct {
	Username string `json:"username" validate:"required,gte=3,lte=50"`
	Password string `json:"password" validate:"required,gte=6,lte=50"`
}
