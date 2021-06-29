package contractions

import "github.com/ZmaximillianZ/stskp_sport_api/internal/models"

// UserRepository is interface to communicate with user storage
type UserRepository interface {
	GetByID(id int) (models.User, error)
	GetByUsername(username string) (models.User, error)
	GetAll() (models.Users, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User) error
}
