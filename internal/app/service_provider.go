package app

import (
	"context"
	"debugger-api/internal/config"
	"debugger-api/internal/controller"
	dataController "debugger-api/internal/controller/data"
	error2 "debugger-api/internal/error"
	messageSource "debugger-api/internal/error/message"
	responseHandler "debugger-api/internal/error/response"
	"debugger-api/internal/repository"
	"debugger-api/internal/repository/data/local"
	redis2 "debugger-api/internal/repository/data/postgres"
	"debugger-api/internal/repository/data/redis"
	"debugger-api/internal/service"
	dataService "debugger-api/internal/service/data"
	"debugger-api/internal/util/properties"
	"fmt"
	redis3 "github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
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
		switch s.config.StorageConfig.Source {
		case config.Local:
			s.dataRepository = local.NewLocalRepository()
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

			s.dataRepository = redis.NewRedisRepository(client)
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

			s.dataRepository = redis2.NewPostgresRepository(pool)
			break
		default:
			panic("incorrect storage type")
		}
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
		messages, err := properties.ReadProperties("./resources/messages.properties")
		if err != nil {
			panic("cannot read messages properties: " + err.Error())
		}

		s.messageSource = messageSource.NewMessageSource(messages)
	}

	return s.messageSource
}
