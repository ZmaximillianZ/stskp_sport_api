package v1

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/routes/api"
	"github.com/labstack/echo/v4"

	"net/http"
)

// GetUsers {
//    \"criteria\": {
//        \"state\": [\">\", \"1\"]
//    },
//    \"limit\": 1,
//    \"offset\": 0,
//    \"order\": {\"id\": \"DESC\"}
//}
// @Summary List users
// @Description Get users with params
// @Produce  json
// @Security JWT
// @Success 200 {array} models.User
// @Failure 500 {object} app.Response
// @Router /api/v1/users [get]
func GetUsers(c echo.Context) error {
	var queryParams api.QueryParams
	if err := c.Bind(&queryParams); err != nil {
		logging.Error(err)
		return c.String(http.StatusNotFound, err.Error())
	}
	// TODO: code after commit is was be repeated very often try export it to middleware
	criteria, order, limit, offset, ok := api.ParseQueryParams(
		models.GetAllowedUserFieldsByMethod("get"),
		&queryParams,
	)
	if !ok {
		return c.String(http.StatusBadGateway, "invalid query params")
	}
	users, err := repository.FindUserBy(criteria, order, limit, offset, models.GetAllowedUserFieldsByMethod("get"))
	if err != nil {
		return c.String(http.StatusBadGateway, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}
