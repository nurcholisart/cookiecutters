package middeware

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func parseReqBody(c *gin.Context) map[string]interface{} {
	var err error

	var bodyByte []byte

	if c.Request.Body == nil {
		bodyByte = []byte("{}")
	} else {
		bodyByte, err = io.ReadAll(c.Request.Body)
		if err != nil {
			bodyByte = []byte("{}")
		}
	}

	body := make(map[string]interface{})
	if err = json.Unmarshal(bodyByte, &body); err != nil {
		body = map[string]interface{}{}
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByte))

	return body
}

func HTTPReqLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqTimestamp := time.Now()

		c.Next()

		reqLatency := time.Since(reqTimestamp)
		logSwitch(c.Writer.Status()).
			Str("request_id", requestid.Get(c)).
			Str("ip_address", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Interface("body", parseReqBody(c)).
			Str("latency", reqLatency.String()).
			Int("status", c.Writer.Status()).
			Msg("http_request")
	}
}

func logSwitch(code int) *zerolog.Event {
	switch {
	case code >= http.StatusInternalServerError:
		return log.Error()
	case code >= http.StatusBadRequest:
		return log.Warn()
	default:
		return log.Info()
	}
}
