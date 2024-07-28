package e2e

import (
	"context"
	"debugger-api/internal/app"
	"debugger-api/internal/config"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"log"
	"testing"
)

type UnitTestSuite struct {
	suite.Suite
	container *postgres.PostgresContainer
	app       *app.App
}

func (s *UnitTestSuite) SetupSuite() {
	ctx := context.Background()

	psql, _ := postgres.RunContainer(
		ctx,
		postgres.WithDatabase("mock"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("password"),
	)
	s.container = psql

	host, _ := s.container.Host(ctx)

	curApp, err := app.NewApp(context.Background(), config.Config{
		StorageConfig: config.Storage{
			Source: "postgres",
			PostgresConfig: config.PostgresConfig{
				Address:  host,
				Password: "password",
				Username: "postgres",
				Database: "mock",
			},
		},
	})
	if err != nil {
		s.T().Fatalf("XD")
	}

	err = curApp.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}

	s.app = curApp
}

func (s *UnitTestSuite) TearDownSuite() {
	_ = s.container.Terminate(context.Background())
	_ = s.app.Stop()
}

func (s *UnitTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Println("before test")
}

func (s *UnitTestSuite) AfterTest(suiteName, testName string) {
	fmt.Println("After test")
}

func (s *UnitTestSuite) TestPost() {
	fmt.Println("XD1")
	require.Error(s.T(), nil)
}

func (s *UnitTestSuite) TestPost2() {
	fmt.Println("XD2")
}

func (s *UnitTestSuite) SubTestPostReturn404() {
	fmt.Println("test1")
}

func (s *UnitTestSuite) SubTestPostReturn200() {
	fmt.Println("test2")
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}
