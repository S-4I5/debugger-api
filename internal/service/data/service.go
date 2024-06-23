package data

import (
	"debugger-api/internal/repository"
	def "debugger-api/internal/service"
)

var _ def.Service = (*service)(nil)

type service struct {
	dataRepository repository.Repository
}

func NewDataService(dataRepository repository.Repository) *service {
	return &service{
		dataRepository: dataRepository,
	}
}
