package controller

import (
	"net/http"
	"strconv"

	"kentang/config"
	"kentang/formatter"
	"kentang/middleware"
	"kentang/models"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	DB := config.Connect()

	user := models.User{}
	c.Bind(&user)

	err := DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, formatter.InternalServerErrorResponse(map[string]interface{}{"error": err.Error()}))
	}

	token, err := middleware.CreateToken(strconv.Itoa(int(user.ID)), user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, formatter.InternalServerErrorResponse(map[string]interface{}{"error": err.Error()}))
	}

	responseData := models.Auth{}
	responseData.ID = user.ID
	responseData.Name = user.Name
	responseData.Email = user.Email
	responseData.Token = token

	return c.JSON(http.StatusOK, formatter.SuccessResponse(responseData))
}
