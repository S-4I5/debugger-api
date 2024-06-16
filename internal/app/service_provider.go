package app

import (
	"debugger-api/internal/config"
	"debugger-api/internal/controller"
	dataController "debugger-api/internal/controller/data"
	error2 "debugger-api/internal/error"
	messageSource "debugger-api/internal/error/message"
	responseHandler "debugger-api/internal/error/response"
	"debugger-api/internal/repository"
	dataRepository "debugger-api/internal/repository/data"
	"debugger-api/internal/service"
	dataService "debugger-api/internal/service/data"
)

type serviceProvider struct {
	dataRepository repository.Repository
	dataService    service.Service
	dataController controller.Controller
	config         config.Config
	errorHandler   error2.Handler
	messageSource  error2.Source
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{
		config: config.MustLoad("./config/config.yaml"),
	}
}

func (s *serviceProvider) DataRepository() repository.Repository {
	if s.dataRepository == nil {
		s.dataRepository = dataRepository.NewRepository()
	}

	return s.dataRepository
}

func (s *serviceProvider) DataService() service.Service {
	if s.dataService == nil {
		s.dataService = dataService.NewDataService(s.DataRepository())
	}

	return s.dataService
}

func (s *serviceProvider) DataController() controller.Controller {
	if s.dataController == nil {
		s.dataController = dataController.NewDataController(s.DataService(), s.ErrorHandler())
	}

	return s.dataController
}

func (s *serviceProvider) ErrorHandler() error2.Handler {
	if s.errorHandler == nil {
		s.errorHandler = responseHandler.NewErrorResponseHandler(s.MessageSource())
	}

	return s.errorHandler
}

func (s *serviceProvider) MessageSource() error2.Source {
	if s.messageSource == nil {
		s.messageSource = messageSource.NewMessageSource()
	}

	return s.messageSource
}
