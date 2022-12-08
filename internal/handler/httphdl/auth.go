package httphdl

import "github.com/labstack/echo/v4"

func (r *httphdl) Auth(c echo.Context) error {
	return c.String(200, "In Auth handler!")
}
