package dto

import "github.com/ZmaximillianZ/stskp_sport_api/internal/models"

const (
	WorkoutTypeCycle           = 1
	WorkoutTypeCompositeParent = 2
	WorkoutTypeComposite       = 3
	WorkoutTypeGymnastic       = 4
	WorkoutTypeSingleCombat    = 5
	WorkoutTypeGame            = 6
)

type WorkoutTypes []WorkoutType

type WorkoutType struct {
	ID       int    `json:"id" db:"id"`
	ParentID *int   `json:"parent_id"`
	Name     string `json:"name" validate:"gte=3,lte=50"`
	Type     uint16 `json:"type" validate:"min=1,max=6"`
}

func LoadWorkoutTypeDTOFromModel(model *models.WorkoutType) *WorkoutType {
	return &WorkoutType{
		ID:       model.ID,
		ParentID: model.ParentID,
		Name:     model.Name,
		Type:     model.Type,
	}
}

func LoadWorkoutTypeModelFromDTO(dto *WorkoutType) *models.WorkoutType {
	return &models.WorkoutType{
		ID:       dto.ID,
		ParentID: dto.ParentID,
		Name:     dto.Name,
		Type:     dto.Type,
	}
}

func LoadWorkoutTypeDTOCollectionFromModel(modelModel *models.WorkoutTypes) *WorkoutTypes {
	var workoutsDTO WorkoutTypes
	for _, workout := range *modelModel {
		workoutModel := workout
		workoutsDTO = append(workoutsDTO, *LoadWorkoutTypeDTOFromModel(&workoutModel))
	}
	return &workoutsDTO
}
