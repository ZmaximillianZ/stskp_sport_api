package controllers

import (
	"net/http"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/dto"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/contractions"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type WorkoutTypeController struct {
	repo         contractions.WorkoutTypeRepository
	errorHandler e.ErrorHandler
	BaseController
}

// NewWorkoutTypeController return new instance of WorkoutTypeController
func NewWorkoutTypeController(
	repo contractions.WorkoutTypeRepository,
	errorHandler e.ErrorHandler,
	v *validator.Validate,
) *WorkoutTypeController {
	return &WorkoutTypeController{
		repo:           repo,
		errorHandler:   errorHandler,
		BaseController: BaseController{*v},
	}
}

// GetByID return workout type by id
// example: /api/v1/workout-type/{id}/
func (ctr *WorkoutTypeController) GetByID(c echo.Context) error {
	var (
		err         error
		workoutType models.WorkoutType
	)
	if workoutType, err = ctr.getWorkoutTypeByID(c); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.LoadWorkoutTypeDTOFromModel(&workoutType))
}

// GetWorkoutTypes return list of workout types
// example: /api/v1/workout-type/all
func (ctr *WorkoutTypeController) GetWorkoutTypes(c echo.Context) error {
	var (
		workoutTypes models.WorkoutTypes
		err          error
	)
	if workoutTypes, err = ctr.repo.GetAll(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.LoadWorkoutTypeDTOCollectionFromModel(&workoutTypes))
}

// Create create workout type
// example: /api/v1/workout-type/create
func (ctr *WorkoutTypeController) Create(c echo.Context) error {
	dtoWorkout := &dto.WorkoutType{}
	if errBindOrValidate := ctr.BindAndValidate(c, dtoWorkout); errBindOrValidate != nil {
		return errBindOrValidate
	}
	if errCreateWorkout := ctr.repo.Create(dto.LoadWorkoutTypeModelFromDTO(dtoWorkout)); errCreateWorkout != nil {
		return errCreateWorkout
	}
	return c.JSON(http.StatusOK, "OK")
}

// Update return workout type by id
// example: /api/v1/workout-type/{id}/
func (ctr *WorkoutTypeController) Update(c echo.Context) error {
	var (
		err         error
		workoutType models.WorkoutType
	)
	if workoutType, err = ctr.getWorkoutTypeByID(c); err != nil {
		return err
	}
	dtoWorkoutType := dto.LoadWorkoutTypeDTOFromModel(&workoutType)
	if errBindOrValidate := ctr.BindAndValidate(c, dtoWorkoutType); errBindOrValidate != nil {
		return errBindOrValidate
	}
	if errUpdateWorkout := ctr.repo.Update(dto.LoadWorkoutTypeModelFromDTO(dtoWorkoutType)); errUpdateWorkout != nil {
		return errUpdateWorkout
	}
	return c.JSON(http.StatusOK, dtoWorkoutType)
}

// Delete return workout type by id
// description: Delete workout by id
// example: /api/v1/workout-type/{id}/ [delete]
func (ctr *WorkoutTypeController) Delete(c echo.Context) error {
	var (
		err         error
		workoutType models.WorkoutType
	)
	if workoutType, err = ctr.getWorkoutTypeByID(c); err != nil {
		return err
	}
	if errDelete := ctr.repo.Delete(&workoutType); errDelete != nil {
		return errDelete
	}
	return c.JSON(http.StatusOK, "OK")
}

func (ctr *WorkoutTypeController) getWorkoutTypeByID(c echo.Context) (models.WorkoutType, error) {
	var (
		id          int64
		err         error
		workoutType models.WorkoutType
	)
	if id, err = ctr.BaseController.GetID(c); err != nil {
		return workoutType, err
	}
	if workoutType, err = ctr.repo.GetByID(int(id)); err != nil {
		return workoutType, err
	}

	return workoutType, err
}
