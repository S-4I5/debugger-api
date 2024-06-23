package data

import (
	error2 "debugger-api/internal/error"
	"debugger-api/internal/service"
)

type controller struct {
	dataService  service.Service
	errorHandler error2.Handler
}

func NewDataController(dataService service.Service, errorHandler error2.Handler) *controller {
	return &controller{
		dataService:  dataService,
		errorHandler: errorHandler,
	}
}
