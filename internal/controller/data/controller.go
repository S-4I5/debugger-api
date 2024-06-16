package data

import (
	error2 "debugger-api/internal/error"
	"debugger-api/internal/service"
)

type Controller struct {
	dataService  service.Service
	errorHandler error2.Handler
}

func NewDataController(dataService service.Service, errorHandler error2.Handler) *Controller {
	return &Controller{
		dataService:  dataService,
		errorHandler: errorHandler,
	}
}
