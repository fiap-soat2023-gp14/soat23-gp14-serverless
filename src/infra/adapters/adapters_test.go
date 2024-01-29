package adapters

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"soat23-gp14_fiap-lambda-application/models"
	"testing"
)

var cfg = aws.Config{}
var client = &cognito.Client{}

type CognitoConfigMock struct {
	mock.Mock
}

type CognitoClientMock struct {
	mock.Mock
}

type AdaptersTestSuite struct {
	suite.Suite
	ctx       context.Context
	userForm  models.UserForm
	userLogin models.UserLogin
	token     string
}

func (s *AdaptersTestSuite) SetupTest() {
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
	s.token = "this-is-a-token"
}

func TestAdaptersTestSuite(t *testing.T) {
	suite.Run(t, new(AdaptersTestSuite))
}

func (c *CognitoConfigMock) LoadDefaultConfig(ctx context.Context, _ ...func(*config.LoadOptions) error) (aws.Config, error) {
	c.Called(ctx)
	if ctx.Value("error") != nil {
		return cfg, errors.New("some-error")
	}
	return cfg, nil
}

func (c *CognitoConfigMock) NewFromConfig(cfg aws.Config, _ ...func(options *cognito.Options)) *cognito.Client {
	c.Called(cfg)
	return client
}

func (c *CognitoClientMock) SignUp(ctx context.Context, params *cognito.SignUpInput, _ ...func(options *cognito.Options)) (*cognito.SignUpOutput, error) {
	args := c.Called(ctx, params)
	return nil, args.Error(1)
}

func (c *CognitoClientMock) InitiateAuth(ctx context.Context, params *cognito.InitiateAuthInput, _ ...func(*cognito.Options)) (*cognito.InitiateAuthOutput, error) {
	args := c.Called(ctx, params)
	return args.Get(0).(*cognito.InitiateAuthOutput), args.Error(1)
}

func (s *AdaptersTestSuite) TestAdapter_SignUp_WhenOk_ShouldRegisterOnProvider() {
	// arrange
	ctx := s.ctx
	c := &CognitoClientMock{}
	c.On("SignUp", ctx, mock.Anything).Return(nil, nil)
	i := New(c)

	// act
	err := i.SignUp(ctx, s.userForm)

	// assert
	s.Nil(err)
}

func (s *AdaptersTestSuite) TestAdapter_SignUp_WhenNotOK_ShouldReturnError() {
	// arrange
	ctx := s.ctx
	c := &CognitoClientMock{}
	c.On("SignUp", ctx, mock.Anything).Return(nil, errors.New("aws-error"))
	i := New(c)

	// act
	err := i.SignUp(ctx, s.userForm)

	// assert
	s.Error(err, "aws-error")
}

func (s *AdaptersTestSuite) TestAdapter_Login_WhenOk_ShouldReturnAccessKey() {
	// arrange
	ctx := s.ctx
	c := &CognitoClientMock{}
	authResponse := &cognito.InitiateAuthOutput{
		AuthenticationResult: &types.AuthenticationResultType{
			AccessToken: &s.token,
		},
	}
	c.On("InitiateAuth", ctx, mock.Anything).Return(authResponse, nil)
	i := New(c)

	// act
	accessToken, err := i.Login(ctx, s.userLogin)

	// assert
	s.Equal(accessToken, s.token)
	s.Nil(err)
}

func (s *AdaptersTestSuite) TestAdapter_Login_WhenError_ShouldReturnError() {
	// arrange
	ctx := s.ctx
	c := &CognitoClientMock{}
	c.On("InitiateAuth", ctx, mock.Anything).Return(&cognito.InitiateAuthOutput{}, errors.New("authentication-error"))
	i := New(c)

	// act
	accessToken, err := i.Login(ctx, s.userLogin)

	// assert
	s.Equal(accessToken, "")
	s.Error(err, "authentication-error")
}

func (s *AdaptersTestSuite) TestAdapter_NewCognitoClient_WhenConfigError_ShouldReturnError() {
	// arrange
	ctx := context.WithValue(context.Background(), "error", "this-is-an-error")
	configMock := &CognitoConfigMock{}
	newConfigFunc = configMock.NewFromConfig
	defaultConfigFunc = configMock.LoadDefaultConfig
	configMock.On("LoadDefaultConfig", ctx).Return(nil, errors.New("some-error"))

	// act
	c, err := NewCognitoClient(ctx)

	// assert
	s.Nil(c)
	s.Error(err, "some-error")
}

func (s *AdaptersTestSuite) TestAdapter_NewCognitoClient_WhenConfigOk_ShouldReturnClient() {
	// arrange
	ctx := s.ctx
	configMock := &CognitoConfigMock{}
	newConfigFunc = configMock.NewFromConfig
	defaultConfigFunc = configMock.LoadDefaultConfig

	configMock.On("NewFromConfig", cfg).Return(client)
	configMock.On("LoadDefaultConfig", ctx).Return(cfg, nil)

	// act
	c, err := NewCognitoClient(ctx)

	// assert
	s.Equal(c, client)
	s.Nil(err)
}
