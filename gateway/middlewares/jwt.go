package middlewares

import (
	"net/http"
	"strings"

	"github.com/bekiryildirimcode/driver-gateway/utils"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")

		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing or invalid token"})
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
		}

		c.Set("user", claims)

		return next(c)
	}
}
