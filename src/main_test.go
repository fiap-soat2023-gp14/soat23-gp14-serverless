package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/aws/aws-lambda-go/events"
)

type mockGinLambda struct {
	mock.Mock
}

func (m *mockGinLambda) ProxyWithContext(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(events.APIGatewayProxyResponse), args.Error(1)
}

type MainTestSuite struct {
	suite.Suite
	mockGinLambda *mockGinLambda
}

func (suite *MainTestSuite) SetupTest() {
	suite.mockGinLambda = new(mockGinLambda)
}

func (suite *MainTestSuite) TestHandler_PingRequest() {
	// arrange
	req := events.APIGatewayProxyRequest{
		Path:       "/ping",
		HTTPMethod: "GET",
	}
	expectedResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{\"message\":\"pong\"}",
	}
	suite.mockGinLambda.On("ProxyWithContext", mock.Anything, req).Return(expectedResponse, nil)

	// act
	actual, err := Handler(context.Background(), req)

	// assert
	suite.NoError(err)
	suite.IsType(actual, events.APIGatewayProxyResponse{})
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
