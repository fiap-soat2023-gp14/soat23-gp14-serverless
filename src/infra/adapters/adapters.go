package adapters

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	awstypes "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/oivinig/soat23-gp14-serverless/infra/settings"
	"github.com/oivinig/soat23-gp14-serverless/models"
)

type Cognito struct {
	client *cognitoidentityprovider.Client
}

var (
	authType = awstypes.AuthFlowTypeUserPasswordAuth
)

func NewIdentityProvider(ctx context.Context) *Cognito {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("unable to load SDK config, %v", err)
	}

	c := cognitoidentityprovider.NewFromConfig(cfg)
	return &Cognito{
		client: c,
	}
}

func (c *Cognito) SignUp(ctx context.Context, u models.UserForm) error {
	_, err := c.client.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
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

func (c *Cognito) Login(ctx context.Context, u models.UserLogin) (string, error) {
	params := map[string]string{
		"PASSWORD": u.Password,
		"USERNAME": u.Username,
	}
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       authType,
		ClientId:       aws.String(settings.GetClientId()),
		AuthParameters: params,
	}

	resp, err := c.client.InitiateAuth(ctx, input)
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
