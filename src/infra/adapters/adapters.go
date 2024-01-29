package adapters

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	awstypes "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/soat-2023-gp14/soat23-gp14_fiap-lambda-application/infra/settings"
	"github.com/soat-2023-gp14/soat23-gp14_fiap-lambda-application/models"
)

var newConfigFunc = cognito.NewFromConfig
var defaultConfigFunc = config.LoadDefaultConfig
var authType = awstypes.AuthFlowTypeUserPasswordAuth

type IdentityProvider struct {
	client Client
}

type Client interface {
	SignUp(ctx context.Context, params *cognito.SignUpInput, optFns ...func(*cognito.Options)) (*cognito.SignUpOutput, error)
	InitiateAuth(ctx context.Context, params *cognito.InitiateAuthInput, optFns ...func(*cognito.Options)) (*cognito.InitiateAuthOutput, error)
}

func New(client Client) *IdentityProvider {
	return &IdentityProvider{
		client: client,
	}
}

func (i *IdentityProvider) SignUp(ctx context.Context, u models.UserForm) error {
	_, err := i.client.SignUp(ctx, &cognito.SignUpInput{
		ClientId: aws.String(settings.GetClientId()),
		Password: aws.String(u.Password),
		Username: aws.String(u.Document),
		UserAttributes: []awstypes.AttributeType{
			*AddAttr("custom:document", u.Document),
		},
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (i *IdentityProvider) Login(ctx context.Context, u models.UserLogin) (string, error) {
	params := map[string]string{
		"PASSWORD": u.Password,
		"USERNAME": u.Username,
	}
	input := &cognito.InitiateAuthInput{
		AuthFlow:       authType,
		ClientId:       aws.String(settings.GetClientId()),
		AuthParameters: params,
	}

	resp, err := i.client.InitiateAuth(ctx, input)
	if err != nil {
		return "", err
	}

	return *resp.AuthenticationResult.AccessToken, nil
}

func AddAttr(name, value string) *awstypes.AttributeType {
	return &awstypes.AttributeType{
		Name:  aws.String(name),
		Value: aws.String(value),
	}
}

func NewCognitoClient(ctx context.Context) (*cognito.Client, error) {
	cfg, err := defaultConfigFunc(ctx)
	if err != nil {
		log.Printf("unable to load SDK config, %v", err)
		return nil, err
	}

	c := newConfigFunc(cfg)
	return c, nil
}
