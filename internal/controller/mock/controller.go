package mock

import (
	error2 "debugger-api/internal/httperr"
	"debugger-api/internal/service"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	dataService  service.MockService
	errorHandler error2.Handler
	validator    validator.Validate
}

func NewController(dataService service.MockService, errorHandler error2.Handler, validate validator.Validate) *controller {
	return &controller{
		validator:    validate,
		dataService:  dataService,
		errorHandler: errorHandler,
	}
}
