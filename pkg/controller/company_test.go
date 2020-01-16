package controller

import (
	"context"
	"testing"

	"github.com/ifaceless/go-starter/pkg/config"

	"github.com/iFaceless/fixture"
	"github.com/stretchr/testify/suite"
)

type SuiteControllerTester struct {
	suite.Suite
	ctx context.Context
	tf  *fixture.TestFixture
}

func (s *SuiteControllerTester) SetupTest() {
	s.tf = config.NewTestFixture()
}

func (s *SuiteControllerTester) TearDownTest() {
	s.tf.DropTables()
}

func (s *SuiteControllerTester) TestGetCompanies() {
	s.tf.Use("company").Test(func() {
		companies, err := GetCompanies()
		s.Nil(err)
		s.Equal(2, len(companies))
	})
}

func TestSuiteController(t *testing.T) {
	suite.Run(t, &SuiteControllerTester{})
}
