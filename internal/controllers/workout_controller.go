package controllers

import (
	"net/http"
	"time"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/dto"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/contractions"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type WorkoutController struct {
	repo         contractions.WorkoutRepository
	errorHandler e.ErrorHandler
	BaseController
}

// NewWorkoutController return new instance of WorkoutController
func NewWorkoutController(repo contractions.WorkoutRepository, errorHandler e.ErrorHandler, v *validator.Validate) *WorkoutController {
	return &WorkoutController{
		repo:           repo,
		errorHandler:   errorHandler,
		BaseController: BaseController{*v},
	}
}

// GetByID return workout by id
// example: /api/v1/workout/{id}/
func (ctr *WorkoutController) GetByID(c echo.Context) error {
	var (
		err     error
		workout models.Workout
	)
	if workout, err = ctr.getWorkoutByID(c); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.LoadWorkoutDTOFromModel(&workout))
}

// GetWorkouts return list of workouts
// example: /api/v1/workout/all
func (ctr *WorkoutController) GetWorkouts(c echo.Context) error {
	var (
		workouts models.Workouts
		err      error
	)
	if workouts, err = ctr.repo.GetAll(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.LoadWorkoutDTOCollectionFromModel(&workouts))
}

// Create create workout
// example: /api/v1/workout/create
func (ctr *WorkoutController) Create(c echo.Context) error {
	var (
		ID  int
		err error
	)
	if ID, err = ctr.BaseController.GetUserIDFromToken(c); err != nil {
		return err
	}
	dtoWorkout := &dto.Workout{UserID: ID, CreatedAt: time.Now()}
	if errBindOrValidate := ctr.BindAndValidate(c, dtoWorkout); errBindOrValidate != nil {
		return errBindOrValidate
	}
	if errCreateWorkout := ctr.repo.Create(dto.LoadWorkoutModelFromDTO(dtoWorkout)); errCreateWorkout != nil {
		return errCreateWorkout
	}
	return c.JSON(http.StatusOK, "OK")
}

// Update return workout by id
// example: /api/v1/workout/{id}/
func (ctr *WorkoutController) Update(c echo.Context) error {
	var (
		err     error
		workout models.Workout
	)
	if workout, err = ctr.getWorkoutByID(c); err != nil {
		return err
	}
	dtoWorkout := dto.LoadWorkoutDTOFromModel(&workout)
	if errBindOrValidate := ctr.BindAndValidate(c, dtoWorkout); errBindOrValidate != nil {
		return errBindOrValidate
	}
	if errUpdateWorkout := ctr.repo.Update(dto.LoadWorkoutModelFromDTO(dtoWorkout)); errUpdateWorkout != nil {
		return errUpdateWorkout
	}
	return c.JSON(http.StatusOK, dtoWorkout)
}

// Delete return workout by id
// description: Delete workout by id
// example: /api/v1/workout/{id}/ [delete]
func (ctr *WorkoutController) Delete(c echo.Context) error {
	var (
		err     error
		workout models.Workout
	)
	if workout, err = ctr.getWorkoutByID(c); err != nil {
		return err
	}
	if errDelete := ctr.repo.Delete(&workout); errDelete != nil {
		return errDelete
	}
	return c.JSON(http.StatusOK, "OK")
}

func (ctr *WorkoutController) getWorkoutByID(c echo.Context) (models.Workout, error) {
	var (
		id      int64
		err     error
		workout models.Workout
	)
	if id, err = ctr.BaseController.GetID(c); err != nil {
		return workout, err
	}
	if workout, err = ctr.repo.GetByID(int(id)); err != nil {
		return workout, err
	}

	return workout, err
}
