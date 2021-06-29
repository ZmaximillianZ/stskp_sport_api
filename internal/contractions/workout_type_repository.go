package contractions

import "github.com/ZmaximillianZ/stskp_sport_api/internal/models"

type WorkoutTypeRepository interface {
	GetByID(id int) (models.WorkoutType, error)
	GetAll() (models.WorkoutTypes, error)
	Create(WorkoutType *models.WorkoutType) error
	Update(WorkoutType *models.WorkoutType) error
	Delete(WorkoutType *models.WorkoutType) error
	// getByIDAndUserID(id, userID int) (models.Workout, error)
}
