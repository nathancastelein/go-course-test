package mocking

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestDatabaseServiceSuite struct {
	suite.Suite
}

func (t *TestDatabaseServiceSuite) SetupSuite() {}

func (t *TestDatabaseServiceSuite) TearDownSuite() {}

func (t *TestDatabaseServiceSuite) SetupTest() {}

func (t *TestDatabaseServiceSuite) TearDownTest() {}

func (t *TestDatabaseServiceSuite) TestListUsers() {}

func (t *TestDatabaseServiceSuite) TestFindUserByID() {}

func (t *TestDatabaseServiceSuite) TestFindUserByIDNoRows() {}

func TestDatabaseService(t *testing.T) {
	suite.Run(t, &TestDatabaseServiceSuite{})
}
