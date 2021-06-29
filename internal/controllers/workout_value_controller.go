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

type WorkoutValueController struct {
	repo         contractions.WorkoutValueRepository
	errorHandler e.ErrorHandler
	BaseController
}

// NewWorkoutValueController return new instance of WorkoutValueController
func NewWorkoutValueController(
	repo contractions.WorkoutValueRepository,
	errorHandler e.ErrorHandler,
	v *validator.Validate,
) *WorkoutValueController {
	return &WorkoutValueController{
		repo:           repo,
		errorHandler:   errorHandler,
		BaseController: BaseController{*v},
	}
}

// GetByID return workout value by id
// example: /api/v1/workout-value/{id}/
func (ctr *WorkoutValueController) GetByID(c echo.Context) error {
	var (
		err          error
		WorkoutValue models.WorkoutValue
	)
	if WorkoutValue, err = ctr.getWorkoutValueByID(c); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.LoadWorkoutValueDTOFromModel(&WorkoutValue))
}

// GetWorkoutValues return list of workout values
// example: /api/v1/workout-value/all
func (ctr *WorkoutValueController) GetWorkoutValues(c echo.Context) error {
	var (
		WorkoutValues models.WorkoutValues
		err           error
	)
	if WorkoutValues, err = ctr.repo.GetAll(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.LoadWorkoutValueDTOCollectionFromModel(&WorkoutValues))
}

// Create create workout value
// example: /api/v1/workout-value/create
func (ctr *WorkoutValueController) Create(c echo.Context) error {
	dtoWorkout := &dto.WorkoutValue{}
	if errBindOrValidate := ctr.BindAndValidate(c, dtoWorkout); errBindOrValidate != nil {
		return errBindOrValidate
	}
	if errCreateWorkout := ctr.repo.Create(dto.LoadWorkoutValueModelFromDTO(dtoWorkout)); errCreateWorkout != nil {
		return errCreateWorkout
	}
	return c.JSON(http.StatusOK, "OK")
}

// Update return workout value by id
// example: /api/v1/workout-value/{id}/
func (ctr *WorkoutValueController) Update(c echo.Context) error {
	var (
		err          error
		WorkoutValue models.WorkoutValue
	)
	if WorkoutValue, err = ctr.getWorkoutValueByID(c); err != nil {
		return err
	}
	dtoWorkoutValue := dto.LoadWorkoutValueDTOFromModel(&WorkoutValue)
	if errBindOrValidate := ctr.BindAndValidate(c, dtoWorkoutValue); errBindOrValidate != nil {
		return errBindOrValidate
	}
	if errUpdateWorkout := ctr.repo.Update(dto.LoadWorkoutValueModelFromDTO(dtoWorkoutValue)); errUpdateWorkout != nil {
		return errUpdateWorkout
	}
	return c.JSON(http.StatusOK, dtoWorkoutValue)
}

// Delete return workout value by id
// description: Delete workout by id
// example: /api/v1/workout-value/{id}/ [delete]
func (ctr *WorkoutValueController) Delete(c echo.Context) error {
	var (
		err          error
		WorkoutValue models.WorkoutValue
	)
	if WorkoutValue, err = ctr.getWorkoutValueByID(c); err != nil {
		return err
	}
	if errDelete := ctr.repo.Delete(&WorkoutValue); errDelete != nil {
		return errDelete
	}
	return c.JSON(http.StatusOK, "OK")
}

func (ctr *WorkoutValueController) getWorkoutValueByID(c echo.Context) (models.WorkoutValue, error) {
	var (
		id           int64
		err          error
		WorkoutValue models.WorkoutValue
	)
	if id, err = ctr.BaseController.GetID(c); err != nil {
		return WorkoutValue, err
	}
	if WorkoutValue, err = ctr.repo.GetByID(int(id)); err != nil {
		return WorkoutValue, err
	}

	return WorkoutValue, err
}
