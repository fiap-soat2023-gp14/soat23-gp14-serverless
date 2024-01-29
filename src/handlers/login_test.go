package handlers

import (
	"bytes"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"net/http"
	"net/http/httptest"
	"testing"

	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var cfg = aws.Config{}
var client = &cognito.Client{}

type CognitoClientMock struct {
	mock.Mock
}

type MockIdentityProvider struct {
	mock.Mock
}

type MockUsersDomain struct {
	mock.Mock
}

type LoginTestSuite struct {
	suite.Suite
}

func (m *CognitoClientMock) NewCognitoClient(ctx context.Context) (*cognito.Client, error) {
	m.Called(ctx)
	return client, nil
}

func TestLoginTestSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}

func (s *LoginTestSuite) TestLogin_WhenBodyIsUnmarshable_ShouldReturnError() {
	// arrange
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{`)))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// act
	Login(c)

	// assert
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

func (s *LoginTestSuite) TestLogin_WhenErrorGeneratingCognitoClient_ShouldReturnError() {
	// arrange
	cc := &CognitoClientMock{}
	cc.On("NewCognitoClient", mock.Anything).Return(nil, errors.New("some-error"))
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{}`)))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// act
	Login(c)

	// assert
	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
}
