package app

import (
	"debugger-api/internal/config"
	controller "debugger-api/internal/controller/data"
	dataController "debugger-api/internal/controller/data/implementation"
	"debugger-api/internal/error/message"
	messageSource "debugger-api/internal/error/message/implementation"
	"debugger-api/internal/error/response"
	responseHandler "debugger-api/internal/error/response/implementaion"
	repository "debugger-api/internal/repository/data"
	dataRepository "debugger-api/internal/repository/data/implementation"
	service "debugger-api/internal/service/data"
	dataService "debugger-api/internal/service/data/implementation"
)

type serviceProvider struct {
	dataRepository repository.Repository
	dataService    service.Service
	dataController controller.Controller
	config         config.Config
	errorHandler   response.Handler
	messageSource  message.Source
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

func (s *serviceProvider) ErrorHandler() response.Handler {
	if s.errorHandler == nil {
		s.errorHandler = responseHandler.NewErrorResponseHandler(s.MessageSource())
	}

	return s.errorHandler
}

func (s *serviceProvider) MessageSource() message.Source {
	if s.messageSource == nil {
		s.messageSource = messageSource.NewMessageSource()
	}

	return s.messageSource
}
