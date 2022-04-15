package middleware

import (
	"demod/lib/jwt"
	"demod/util/xrsp"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")
		if token == "" {
			token = c.QueryParam("token")
		}

		tokenData, err := jwt.ParseToken(token)
		if err != nil || tokenData.UserId == 0 || tokenData.ExpireAt < time.Now().Unix() {
			return c.JSON(http.StatusUnauthorized, xrsp.ErrorText("invalid token"))
		}

		c.Set("user_id", tokenData.UserId)

		return next(c)
	}
}
