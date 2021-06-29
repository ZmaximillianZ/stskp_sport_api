package dto

import (
	"time"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
)

const (
	UnitMeters      = 1
	UnitSeconds     = 2
	UnitKilograms   = 3
	UnitReiteration = 4
	UnitWatt        = 5
	UnitTemp        = 6
	UnitPulse       = 7
)

type WorkoutValues []WorkoutValue

type WorkoutValue struct {
	ID            int        `json:"id" db:"id"`
	WorkoutID     int        `json:"workout_id"`
	WorkoutTypeID int        `json:"workout_type_id"`
	Value         *float64   `json:"value"`
	Unit          int8       `json:"unit" validate:"min=1,max=7"`
	StartedAt     *time.Time `json:"started_at"`
	EndedAt       *time.Time `json:"ended_at"`
}

func LoadWorkoutValueDTOFromModel(model *models.WorkoutValue) *WorkoutValue {
	return &WorkoutValue{
		ID:            model.ID,
		WorkoutID:     model.WorkoutID,
		WorkoutTypeID: model.WorkoutTypeID,
		Value:         model.Value,
		Unit:          model.Unit,
		StartedAt:     model.StartedAt,
		EndedAt:       model.EndedAt,
	}
}

func LoadWorkoutValueModelFromDTO(dto *WorkoutValue) *models.WorkoutValue {
	return &models.WorkoutValue{
		ID:            dto.ID,
		WorkoutID:     dto.WorkoutID,
		WorkoutTypeID: dto.WorkoutTypeID,
		Value:         dto.Value,
		Unit:          dto.Unit,
		StartedAt:     dto.StartedAt,
		EndedAt:       dto.EndedAt,
	}
}

func LoadWorkoutValueDTOCollectionFromModel(workoutValues *models.WorkoutValues) *WorkoutValues {
	var workoutValuesDTO WorkoutValues
	for _, workout := range *workoutValues {
		workoutModel := workout
		workoutValuesDTO = append(workoutValuesDTO, *LoadWorkoutValueDTOFromModel(&workoutModel))
	}
	return &workoutValuesDTO
}
