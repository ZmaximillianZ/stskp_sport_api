package controllers

import (
	"net/http"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/contractions"
	"github.com/gin-gonic/gin"
)

// UserController is HTTP controller for manage users
type UserController struct {
	repo contractions.UserRepository
}

// NewUserController return new instance of UserController
func NewUserController(repo contractions.UserRepository) *UserController {
	return &UserController{repo}
}

// GetUserByID return user by id
// @Summary Show a user
// @Description Get user by id
// @Produce  json
// @Security JWT
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Header 200 {string} X-AUTH-TOKEN "qwerty"
// @Failure 500 {object} app.Response
// @Router /api/v1/users/{id}/ [get]
func (ctr *UserController) GetUserByID(c *gin.Context) {
	user, err := ctr.repo.GetByID(c.GetInt("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}