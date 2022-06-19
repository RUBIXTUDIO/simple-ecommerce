package controllers

import (
	"api/blueprints/users/databases"
	"api/blueprints/users/models"
	"net/http"

	"github.com/labstack/echo"
)

func UserRegister(c echo.Context) error {
	//get user's input
	user := models.Users{}
	c.Bind(&user)

	//check is email exists?
	// is_email_exists, _ := databases.EmailCheck(user.Email)
	// if is_email_exists != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "Email has already exist",
	// 	})
	// }

	//create new user
	new_user, err := databases.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	type Output struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	//set output data
	output := Output{
		ID:    new_user.ID,
		Email: new_user.Email,
		Name:  new_user.Name,
	}

	return c.JSON(http.StatusOK, output)
}

func GetAllUsers(c echo.Context) error {
	users, err := databases.ReadAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
