package protocol

import (
	"playground/internal/handler/httphdl"
	"playground/internal/handler/middleware"

	"github.com/labstack/echo/v4"
)

func ServeREST() {
	e := echo.New()
	custService := app.svr
	custHttp := httphdl.NewHttphdl(custService)
	e.POST("/", custHttp.CreateCustomer)
	e.PUT("/:id", custHttp.UpdateCustomer)
	e.DELETE("/:id", custHttp.DeleteCustomer)
	e.GET("/", custHttp.ListCustomer)
	e.POST("/login", custHttp.Login)
	e.GET("/auth", custHttp.Auth, middleware.RequireAuth)

	e.Start(":8000")
}
