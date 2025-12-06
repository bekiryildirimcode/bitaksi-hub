package handlers

import (
	"net/http"

	"github.com/bekiryildirimcode/driver-gateway/utils"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(c echo.Context) error {
	req := new(LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	token, err := utils.GenerateToken(req.FirstName, req.LastName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Token creation failed"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
