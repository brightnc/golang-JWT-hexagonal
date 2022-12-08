package middleware

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token, err := c.Cookie("Authorization")
		if err != nil {
			return c.String(401, "UnAuthorization")
		}
		tokenStr := token.Value

		tokenJwt, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			fmt.Println(err)
			return c.String(401, "UnAuthorization")
		}

		if !tokenJwt.Valid {
			return c.String(401, "Token is not valid")
		}
		return next(c)

	}
}
