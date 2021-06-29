package dto

import (
	"time"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
)

type Users []User

type User struct {
	ID        int       `json:"id" db:"id" validate:"required"`
	Username  string    `json:"username"   validate:"gte=3,lte=50" `
	Password  string    `json:"password"   ` //validate:"gte=6,lte=50"`
	State     int8      `json:"state"      validate:"min=1,max=5"`
	CreatedAt time.Time `json:"created_at"`
	Email     *string   `json:"email"      validate:"email"`
}

func LoadUserDTOFromModel(model *models.User) *User {
	return &User{
		ID:        model.ID,
		Username:  model.Username,
		Password:  model.Password,
		State:     model.State,
		CreatedAt: model.CreatedAt,
		Email:     model.Email,
	}
}

func LoadUserModelFromDTO(dto *User) *models.User {
	return &models.User{
		ID:        dto.ID,
		Username:  dto.Username,
		Password:  dto.Password,
		State:     dto.State,
		CreatedAt: dto.CreatedAt,
		Email:     dto.Email,
	}
}

func LoadUserDTOCollectionFromModel(usersModel *models.Users) *Users {
	var usersDTO Users
	for _, user := range *usersModel {
		userModel := user
		usersDTO = append(usersDTO, *LoadUserDTOFromModel(&userModel))
	}
	return &usersDTO
}
