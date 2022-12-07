package protocol

import (
	"playground/internal/handler/httphdl"

	"github.com/labstack/echo/v4"
)

func ServeREST() {
	e := echo.New()
	custService := app.svr
	custHttp := httphdl.NewHttphdl(custService)

	e.GET("/", custHttp.GetAllCustomers)
	e.GET("/:id", custHttp.GetCustomerByID)
	e.POST("/", custHttp.CreateCustomer)

	e.Start(":8000")
}
