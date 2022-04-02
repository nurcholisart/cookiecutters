package api

import (
	"{{cookiecutter.module_name}}/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	Conf config.Config
}

func NewAPI(conf *config.Config) API {
	return API{
		Conf: *conf,
	}
}

func (a *API) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (a *API) RouteNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "resource not found",
	})
}

func (a *API) AbortWebhookHandling(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"ok":      false,
		"message": message,
	})
}
