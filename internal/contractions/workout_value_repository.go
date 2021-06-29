package contractions

import "github.com/ZmaximillianZ/stskp_sport_api/internal/models"

type WorkoutValueRepository interface {
	GetByID(id int) (models.WorkoutValue, error)
	GetAll() (models.WorkoutValues, error)
	Create(WorkoutValue *models.WorkoutValue) error
	Update(WorkoutValue *models.WorkoutValue) error
	Delete(WorkoutValue *models.WorkoutValue) error
}
