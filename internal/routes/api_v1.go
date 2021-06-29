package routes

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RegisterAPIV1 initialize routing information
func RegisterAPIV1(
	router *echo.Group,
	userController *controllers.UserController,
	workoutController *controllers.WorkoutController,
	workoutTypeController *controllers.WorkoutTypeController,
	workoutValueController *controllers.WorkoutValueController,
) {
	jwt := middleware.JWT([]byte("get_key_from_env"))

	router.POST("/auth", userController.Authenticate)
	router.POST("/create", userController.Create)
	user := router.Group("/users")
	user.GET("/", userController.GetUsers, jwt)
	user.POST("/", userController.Create)
	user.GET("/:id", userController.GetByID, jwt)
	user.PUT("/:id", userController.Update, jwt)
	user.DELETE("/:id", userController.Delete, jwt)

	workout := router.Group("/workout", jwt)
	workout.GET("/all", workoutController.GetWorkouts)
	workout.POST("/create", workoutController.Create)
	workout.GET("/:id", workoutController.GetByID)
	workout.PUT("/:id", workoutController.Update)
	workout.DELETE("/:id", workoutController.Delete)

	workoutType := router.Group("/workout-type", jwt)
	workoutType.GET("/all", workoutTypeController.GetWorkoutTypes)
	workoutType.POST("/create", workoutTypeController.Create)
	workoutType.GET("/:id", workoutTypeController.GetByID)
	workoutType.PUT("/:id", workoutTypeController.Update)
	workoutType.DELETE("/:id", workoutTypeController.Delete)

	workoutValue := router.Group("/workout-value", jwt)
	workoutValue.GET("/all", workoutValueController.GetWorkoutValues)
	workoutValue.POST("/create", workoutValueController.Create)
	workoutValue.GET("/:id", workoutValueController.GetByID)
	workoutValue.PUT("/:id", workoutValueController.Update)
	workoutValue.DELETE("/:id", workoutValueController.Delete)
}
