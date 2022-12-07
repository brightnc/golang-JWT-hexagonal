package protocol

import (
	"playground/infrastructure"
	"playground/internal/core/service"
	"playground/internal/repository"
)

var app *application

type application struct {
	svr *service.CustomerService
}

func init() {
	db := infrastructure.InitDatabase()
	custRepository := repository.NewCustomerRepositoryDB(db)
	custService := service.NewCustomerService(&custRepository)
	app = &application{
		svr: &custService,
	}
}
