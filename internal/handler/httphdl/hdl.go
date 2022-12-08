package httphdl

import (
	"encoding/json"
	"fmt"
	"playground/internal/core/domain"
	"playground/internal/core/port"
	"playground/internal/handler/httphdl/dto"
	"strconv"

	"github.com/labstack/echo/v4"
)

type httphdl struct {
	srv port.CustomerService
}

func NewHttphdl(srv port.CustomerService) *httphdl {
	return &httphdl{srv: srv}
}

// func (r *httphdl) GetAllCustomers(c echo.Context) error {
// 	customers, err := r.srv.GetAllCustomer()
// 	if err != nil {
// 		fmt.Println(err)
// 		return echo.ErrBadRequest
// 	}
// 	return c.JSON(200, customers)
// }

// func (r *httphdl) GetCustomerByID(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	customer, err := r.srv.GetCustomerByID(id)
// 	if err != nil {
// 		fmt.Println(err)
// 		return echo.ErrBadRequest
// 	}

// 	return c.JSON(200, customer)
// }

func (r *httphdl) CreateCustomer(c echo.Context) error {
	var request *dto.Customer
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	customer := request.ToDomainCustomer()
	custRes, err := r.srv.CreateUser(&customer)
	if err != nil {
		fmt.Println(err)
	}

	custResponse := domain.CustomerResponse{
		Username:  custRes.Username,
		Email:     custRes.Email,
		CreatedAt: custRes.CreatedAt,
	}

	return c.JSON(201, custResponse)
}

func (r *httphdl) UpdateCustomer(c echo.Context) error {
	var request *dto.Customer
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}

	customer := request.ToDomainCustomer()
	custRes, err := r.srv.UpdateUser(idInt, &customer)
	if err != nil {
		fmt.Println(err)
	}

	custResponse := domain.CustomerResponse{
		Username: custRes.Username,
		Email:    custRes.Email,
	}

	return c.JSON(200, custResponse)
}

func (r *httphdl) DeleteCustomer(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	err = r.srv.DeleteUser(idInt)
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(200, "success")
}

func (r *httphdl) ListCustomer(c echo.Context) error {
	customers, err := r.srv.ListUsers()
	if err != nil {
		fmt.Println(err)
		return echo.ErrBadRequest
	}
	return c.JSON(200, customers)
}
