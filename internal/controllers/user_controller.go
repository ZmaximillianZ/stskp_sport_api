package controllers

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/routers/api"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/utils"
	"github.com/astaxie/beego/validation"

	"net/http"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/app"
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

// @Summary Get Auth
// @Description user authorization
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Header 200 {string} Access-Control-Allow_origin "*"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/user/auth [post]
func (ctr *UserController) GetAuth(c *gin.Context) {
	valid := validation.Validation{}

	username, _ := c.GetQuery("username")
	password, _ := c.GetQuery("password")
	a := api.Auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		if valid.HasErrors() {
			// maybe c.Error(valid.Errors)
			app.MarkErrors(valid.Errors)
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user, err := ctr.repo.GetByUsername(username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	invalid := user.InvalidPassword(password)

	if invalid {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid password")
		return
	}

	token, err := utils.GenerateToken(a.Username, a.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"token": token})
}
