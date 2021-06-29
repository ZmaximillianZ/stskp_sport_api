package contractions

import "github.com/ZmaximillianZ/stskp_sport_api/internal/models"

type WorkoutRepository interface {
	GetByID(id int) (models.Workout, error)
	GetAll() (models.Workouts, error)
	Create(Workout *models.Workout) error
	Update(Workout *models.Workout) error
	Delete(Workout *models.Workout) error
	// getByIDAndUserID(id, userID int) (models.Workout, error)
}
