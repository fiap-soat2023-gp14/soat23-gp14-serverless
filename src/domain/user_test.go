package domain

import (
	"context"
	"errors"
	"github.com/oivinig/soat23-gp14-serverless/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuthMock struct {
	mock.Mock
}

func (a *AuthMock) SignUp(ctx context.Context, u models.UserForm) error {
	args := a.Called(ctx, u)
	return args.Error(0)
}

func (a *AuthMock) Login(ctx context.Context, u models.UserLogin) (string, error) {
	args := a.Called(ctx, u)
	return args.String(0), args.Error(1)
}

type UserTestSuite struct {
	suite.Suite
	ctx       context.Context
	userForm  models.UserForm
	userLogin models.UserLogin
}

func (s *UserTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.userForm = models.UserForm{
		Name:     "Tom Brady",
		Document: "15600204078",
		Email:    "tom.brady@email.com",
		Password: "this-is-a-very-safe-password",
	}
	s.userLogin = models.UserLogin{
		Username: "15600204078",
		Password: "this-is-a-very-safe-password",
	}
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) TestUser_CreateUser_WhenAuthReturnsError_ThenShouldBeError() {
	// arrange
	ctx := s.ctx
	usr := s.userForm
	provider := new(AuthMock)
	provider.On("SignUp", ctx, mock.Anything).Return(errors.New("error-during-test"))

	// act
	u := NewUsersDomain(provider)
	err := u.CreateUser(ctx, usr)

	// assert
	s.Error(err, "error-during-test")
}

func (s *UserTestSuite) TestUser_CreateUser_WhenAuthOK_ThenShouldBeNil() {
	// arrange
	ctx := s.ctx
	usr := s.userForm
	provider := new(AuthMock)
	provider.On("SignUp", ctx, mock.Anything).Return(nil)

	// act
	u := NewUsersDomain(provider)
	err := u.CreateUser(ctx, usr)

	// assert
	s.Nil(err)
}

func (s *UserTestSuite) TestUser_CreateUser_WhenDocumentInvalid_ThenShouldBeError() {
	// arrange
	ctx := s.ctx
	usr := s.userForm
	usr.Document = "111222333-00"
	provider := new(AuthMock)

	// act
	u := NewUsersDomain(provider)
	err := u.CreateUser(ctx, usr)

	// assert
	s.Error(err, ErrInvalidDocument)
}

func (s *UserTestSuite) TestUser_CreateUser_WhenDocumentLengthInvalid_ThenShouldBeError() {
	// arrange
	ctx := s.ctx
	usr := s.userForm
	usr.Document = "111222"
	provider := new(AuthMock)

	// act
	u := NewUsersDomain(provider)
	err := u.CreateUser(ctx, usr)

	// assert
	s.Error(err, ErrInvalidDocument)
}

func (s *UserTestSuite) TestUser_Login_WhenDocumentLengthInvalid_ThenShouldBeError() {
	// arrange
	ctx := s.ctx
	usr := s.userLogin
	usr.Username = ""
	provider := new(AuthMock)

	// act
	u := NewUsersDomain(provider)
	_, err := u.Login(ctx, usr)

	// assert
	s.Error(err, ErrInvalidDocument)
}

func (s *UserTestSuite) TestUser_Login_WhenAuthHasError_ThenShouldBeError() {
	// arrange
	ctx := s.ctx
	usr := s.userLogin
	provider := new(AuthMock)
	provider.On("Login", ctx, mock.Anything).Return("", errors.New("this is an generic error"))

	// act
	u := NewUsersDomain(provider)
	_, err := u.Login(ctx, usr)

	// assert
	s.Error(err, "this is an generic error")
}

func (s *UserTestSuite) TestUser_Login_WhenAuthOK_ThenShouldReturnAccessToken() {
	// arrange
	ctx := s.ctx
	usr := s.userLogin
	provider := new(AuthMock)
	provider.On("Login", ctx, mock.Anything).Return("access-token", nil)

	// act
	u := NewUsersDomain(provider)
	accessToken, err := u.Login(ctx, usr)

	// assert
	s.Nil(err)
	s.Equal(accessToken, "access-token")
}
