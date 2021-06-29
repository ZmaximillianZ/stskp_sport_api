package dto

import (
	"time"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
)

type Workouts []Workout

type Workout struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description" validate:"required,gte=5,lte=255"`
}

func LoadWorkoutDTOFromModel(model *models.Workout) *Workout {
	return &Workout{
		ID:          model.ID,
		UserID:      model.UserID,
		CreatedAt:   model.CreatedAt,
		Description: model.Description,
	}
}

func LoadWorkoutModelFromDTO(dto *Workout) *models.Workout {
	return &models.Workout{
		ID:          dto.ID,
		UserID:      dto.UserID,
		CreatedAt:   dto.CreatedAt,
		Description: dto.Description,
	}
}

func LoadWorkoutDTOCollectionFromModel(modelModel *models.Workouts) *Workouts {
	var workoutsDTO Workouts
	for _, workout := range *modelModel {
		workoutModel := workout
		workoutsDTO = append(workoutsDTO, *LoadWorkoutDTOFromModel(&workoutModel))
	}
	return &workoutsDTO
}
