package router

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"{{cookiecutter.module_name}}/api"
	"{{cookiecutter.module_name}}/config"
	"{{cookiecutter.module_name}}/middeware"
)

func NewRouter(conf *config.Config) *gin.Engine {
	r := gin.New()

	r.Use(requestid.New())
	r.Use(middeware.HTTPReqLog())

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			log.Error().Msg(err)
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong on our side",
		})

		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	httpAPI := api.NewAPI(conf)

	r.GET("/", httpAPI.Home)
	r.GET("/internal/healthcheck", httpAPI.Healthcheck)

	r.NoRoute(httpAPI.RouteNotFound)

	return r
}
