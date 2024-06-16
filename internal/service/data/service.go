package data

import (
	"debugger-api/internal/repository"
	def "debugger-api/internal/service"
)

var _ def.Service = (*Service)(nil)

type Service struct {
	dataRepository repository.Repository
}

func NewDataService(dataRepository repository.Repository) *Service {
	return &Service{
		dataRepository: dataRepository,
	}
}
