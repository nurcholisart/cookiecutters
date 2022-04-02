package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"services": gin.H{
			"qiscus-sdk": "ok",
			"qismo":      "ok",
		},
		"ok": true,
	})
}
