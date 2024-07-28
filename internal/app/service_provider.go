package app

import (
	"context"
	"debugger-api/internal/config"
	"debugger-api/internal/controller"
	dataController "debugger-api/internal/controller/mock"
	"debugger-api/internal/httperr"
	error2 "debugger-api/internal/httperr"
	"debugger-api/internal/middleware"
	"debugger-api/internal/middleware/compressor"
	"debugger-api/internal/middleware/fullurl"
	"debugger-api/internal/middleware/requestid"
	"debugger-api/internal/repository"
	"debugger-api/internal/repository/mock/local"
	redis2 "debugger-api/internal/repository/mock/postgres"
	"debugger-api/internal/repository/mock/redis"
	"debugger-api/internal/service"
	dataService "debugger-api/internal/service/mock"
	"debugger-api/internal/util/properties"
	"fmt"
	"github.com/go-playground/validator/v10"
	redis3 "github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

type serviceProvider struct {
	dataRepository                   repository.MockRepository
	dataService                      service.MockService
	dataController                   controller.MockController
	requestIdMiddlewareProvider      middleware.RequestIdMiddlewareProvider
	fullRequestUrlMiddlewareProvider middleware.FullRequestUrlMiddlewareProvider
	compressorMiddlewareProvider     middleware.CompressorMiddlewareProvider
	config                           config.Config
	errorHandler                     error2.Handler
	messageSource                    error2.Source
}

func newServiceProvider(cfg config.Config) *serviceProvider {
	return &serviceProvider{
		config: cfg,
	}
}

func (s *serviceProvider) CompressorMiddlewareProvider() middleware.CompressorMiddlewareProvider {
	if s.compressorMiddlewareProvider == nil {
		s.compressorMiddlewareProvider = compressor.NewMiddlewareProvider()
	}
	return s.compressorMiddlewareProvider
}

func (s *serviceProvider) FullRequestUrlMiddlewareProvider() middleware.FullRequestUrlMiddlewareProvider {
	if s.fullRequestUrlMiddlewareProvider == nil {
		s.fullRequestUrlMiddlewareProvider = fullurl.NewMiddlewareProvider()
	}
	return s.fullRequestUrlMiddlewareProvider
}

func (s *serviceProvider) RequestIdMiddlewareProvider() middleware.RequestIdMiddlewareProvider {
	if s.requestIdMiddlewareProvider == nil {
		s.requestIdMiddlewareProvider = requestid.NewMiddlewareProvider()
	}
	return s.requestIdMiddlewareProvider
}

func (s *serviceProvider) DataRepository() repository.MockRepository {
	if s.dataRepository == nil {
		switch s.config.StorageConfig.Source {
		case config.Local:
			s.dataRepository = local.NewRepository()
			break
		case config.Redis:
			redisConfig := &redis3.Options{
				Addr: s.config.StorageConfig.RedisConfig.Address,
				DB:   s.config.StorageConfig.RedisConfig.Database,
			}

			if s.config.StorageConfig.RedisConfig.Username != "" {
				redisConfig.Username = s.config.StorageConfig.RedisConfig.Username
			}

			if s.config.StorageConfig.RedisConfig.Password != "" {
				redisConfig.Password = s.config.StorageConfig.RedisConfig.Password
			}

			client := redis3.NewClient(redisConfig)

			s.dataRepository = redis.NewRepository(client)
			break
		case config.Postgres:
			pool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s/%s",
				s.config.StorageConfig.PostgresConfig.Username,
				s.config.StorageConfig.PostgresConfig.Password,
				s.config.StorageConfig.PostgresConfig.Address,
				s.config.StorageConfig.PostgresConfig.Database,
			))
			if err != nil {
				panic("cannot create pgx pool")
			}

			s.dataRepository = redis2.NewRepository(pool)
			break
		default:
			panic("incorrect storage type")
		}
	}

	return s.dataRepository
}

func (s *serviceProvider) DataService() service.MockService {
	if s.dataService == nil {
		s.dataService = dataService.NewDataService(s.DataRepository())
	}

	return s.dataService
}

func (s *serviceProvider) DataController() controller.MockController {
	if s.dataController == nil {
		s.dataController = dataController.NewController(s.DataService(), s.ErrorHandler(), *validator.New(validator.WithRequiredStructEnabled()))
	}
	return s.dataController
}

func (s *serviceProvider) ErrorHandler() error2.Handler {
	if s.errorHandler == nil {
		s.errorHandler = httperr.NewErrorResponseHandler(s.MessageSource())
	}

	return s.errorHandler
}

func (s *serviceProvider) MessageSource() error2.Source {
	if s.messageSource == nil {
		messages, err := properties.ReadProperties("./resources/messages.properties")
		if err != nil {
			panic("cannot read messages properties: " + err.Error())
		}

		s.messageSource = httperr.NewMessageSource(messages)
	}

	return s.messageSource
}
