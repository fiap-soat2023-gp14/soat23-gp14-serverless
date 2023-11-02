package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oivinig/soat23-gp14-serverless/domain"
	"github.com/oivinig/soat23-gp14-serverless/infra/adapters"
	"github.com/oivinig/soat23-gp14-serverless/models"
)

func Login(c *gin.Context) {
	ctx := &gin.Context{}
	body := models.UserLogin{}
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
	accessToken, err := domain.Login(ctx, body)
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
