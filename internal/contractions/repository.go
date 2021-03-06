package contractions

import "github.com/ZmaximillianZ/stskp_sport_api/internal/models"

// UserRepository is interface to communicate with user storage
type UserRepository interface {
	GetByID(id int) (models.User, error)
	GetByUsername(username string) (models.User, error)
	GetUsers() (models.Users, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}
