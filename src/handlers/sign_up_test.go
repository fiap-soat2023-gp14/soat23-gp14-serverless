package handlers

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SignUpTestSuite struct {
	suite.Suite
}

func TestSignUpTestSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}

func (s *LoginTestSuite) TestSignUp_WhenBodyIsUnmarshable_ShouldReturnError() {
	// arrange
	req, _ := http.NewRequest("POST", "/sign-up", bytes.NewBuffer([]byte(`{`)))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// act
	SignUp(c)

	// assert
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

func (s *LoginTestSuite) TestSignUp_WhenErrorGeneratingCognitoClient_ShouldReturnError() {
	// arrange
	cc := &CognitoClientMock{}
	cc.On("NewCognitoClient", mock.Anything).Return(nil, errors.New("some-error"))
	req, _ := http.NewRequest("POST", "/sign-up", bytes.NewBuffer([]byte(`{}`)))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// act
	SignUp(c)

	// assert
	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
}
