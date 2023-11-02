package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oivinig/soat23-gp14-serverless/domain"
	"github.com/oivinig/soat23-gp14-serverless/infra/adapters"
	"github.com/oivinig/soat23-gp14-serverless/models"
)

func SignUp(c *gin.Context) {
	ctx := &gin.Context{}
	body := models.UserForm{}
	err := c.BindJSON(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	identityProvider := adapters.NewIdentityProvider(ctx)
	domain := domain.NewUsersDomain(identityProvider)
	if err := domain.CreateUser(ctx, body); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "user created",
	})
}
