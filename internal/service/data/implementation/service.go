package implementation

import (
	"debugger-api/internal/repository/data"
	def "debugger-api/internal/service/data"
)

var _ def.Service = (*Service)(nil)

type Service struct {
	dataRepository data.Repository
}

func NewDataService(dataRepository data.Repository) *Service {
	return &Service{
		dataRepository: dataRepository,
	}
}
