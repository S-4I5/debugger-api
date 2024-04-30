package implementation

import (
	"debugger-api/internal/error/response"
	"debugger-api/internal/service/data"
)

type Controller struct {
	dataService  data.Service
	errorHandler response.Handler
}

func NewDataController(dataService data.Service, errorHandler response.Handler) *Controller {
	return &Controller{
		dataService:  dataService,
		errorHandler: errorHandler,
	}
}
