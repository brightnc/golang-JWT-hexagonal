package httphdl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"playground/internal/handler/httphdl/dto"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (r *httphdl) Login(c echo.Context) error {
	var req *dto.Login
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		panic(err)
	}
	fmt.Println(req.Email)
	customer, isExist := r.srv.FindUser(req.Email)
	fmt.Println(isExist)
	if !isExist {
		return echo.ErrNotFound
	}

	passwordHashReq := []byte(req.Password)
	passwordHashDb := []byte(customer.Password)

	err = bcrypt.CompareHashAndPassword(passwordHashDb, passwordHashReq)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, "invalid credentials")
	}
	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": customer.Id,
		"exp": now.Add(time.Hour * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.JSON(400, "Failed to create Token!")
	}
	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = tokenString
	cookie.Expires = now.Add(time.Hour * 24 * 30)
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	c.SetCookie(cookie)

	return c.String(http.StatusOK, "write a cookie")
}
