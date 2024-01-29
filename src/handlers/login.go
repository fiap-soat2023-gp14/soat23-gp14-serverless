package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oivinig/soat23-gp14-serverless/domain"
	"github.com/oivinig/soat23-gp14-serverless/infra/adapters"
	"github.com/oivinig/soat23-gp14-serverless/models"
)

var newCognitoClientFunc = adapters.NewCognitoClient
var newIdentityProviderFunc = adapters.New
var newUsersDomainFunc = domain.NewUsersDomain

func Login(c *gin.Context) {
	ctx := &gin.Context{}
	body := models.UserLogin{}
	err := c.BindJSON(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	cognitoClient, err := newCognitoClientFunc(ctx)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	identityProvider := newIdentityProviderFunc(cognitoClient)
	d := newUsersDomainFunc(identityProvider)
	accessToken, err := d.Login(ctx, body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.Header("Access-Token", accessToken)
	c.JSON(http.StatusAccepted, gin.H{
		"message": "success",
	})
}
