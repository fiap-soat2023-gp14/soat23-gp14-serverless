// settings_test.go
package settings

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

func setEnvVar(key, value string) {
	_ = os.Setenv(key, value)
}

type EnvTestSuite struct {
	suite.Suite
}

func TestEnvTestSuiteSuite(t *testing.T) {
	suite.Run(t, new(EnvTestSuite))
}

func (s *EnvTestSuite) TestGetUserPoolId_WhenValueIsSet_ShouldReturnValue() {
	// arrange
	value := "test_user_pool_id"
	setEnvVar("AWS_USER_POOL_ID", value)

	// act
	actual := GetUserPoolId()

	// assert
	s.Equal(actual, value)
}

func (s *EnvTestSuite) TestGetRegion_WhenValueIsSet_ShouldReturnValue() {
	// arrange
	value := "us-east-1"
	setEnvVar("AWS_USER_POOL_REGION", value)

	// act
	actual := GetRegion()

	// assert
	s.Equal(actual, value)
}

func (s *EnvTestSuite) TestGetClientId_WhenValueIsSet_ShouldReturnValue() {
	// arrange
	value := "test_client_id"
	setEnvVar("AWS_CLIENT_ID", value)

	// act
	actual := GetClientId()

	// assert
	s.Equal(actual, value)
}
