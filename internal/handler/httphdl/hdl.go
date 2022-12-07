package httphdl

import (
	"encoding/json"
	"fmt"
	"playground/internal/core/domain"
	"playground/internal/core/port"
	"playground/internal/handler/httphdl/dto"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type httphdl struct {
	srv port.CustomerService
}

func NewHttphdl(srv port.CustomerService) *httphdl {
	return &httphdl{srv: srv}
}

func (r *httphdl) GetAllCustomers(c echo.Context) error {
	customers, err := r.srv.GetAllCustomer()
	if err != nil {
		fmt.Println(err)
		return echo.ErrBadRequest
	}
	return c.JSON(200, customers)
}

func (r *httphdl) GetCustomerByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := r.srv.GetCustomerByID(id)
	if err != nil {
		fmt.Println(err)
		return echo.ErrBadRequest
	}

	return c.JSON(200, customer)
}

func (r *httphdl) CreateCustomer(c echo.Context) error {
	var request *dto.Customer
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	hasedByte, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		fmt.Println(err)
	}
	request.Password = string(hasedByte)
	customer := request.ToDomainCustomer()
	custRes, err := r.srv.CreateCustomer(&customer)
	if err != nil {
		fmt.Println(err)
	}

	custResponse := domain.CustomerResponse{
		CustomerID: custRes.CustomerID,
		Username:   custRes.Username,
		Email:      custRes.Email,
	}

	return c.JSON(201, custResponse)
}
