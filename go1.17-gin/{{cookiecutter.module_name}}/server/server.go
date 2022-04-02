package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"{{cookiecutter.module_name}}/config"
	"{{cookiecutter.module_name}}/helpers"
	"syscall"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Start(conf *config.Config) {
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := setupRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	helpers.PrintInfoStringLog("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		helpers.PrintFatalStringLog(fmt.Sprintf("Server forced to shutdown: %s", err.Error()))
	}

	helpers.PrintInfoStringLog("Server exitting")
}

func setupRouter() *gin.Engine {
	r := gin.New()

	r.Use(requestid.New())
	r.Use(customLogMiddleware())

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			helpers.PrintErrStringLog(err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wrong on our side",
			})
		}

		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})

	r.GET("/internal/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})

	return r
}

type httpRequestLog struct {
	RequestID  string                 `json:"request_id"`
	Timestamp  int64                  `json:"timestamp,omitempty"`
	Method     string                 `json:"method,omitempty"`
	Path       string                 `json:"path,omitempty"`
	StatusCode int                    `json:"status,omitempty"`
	IPAddress  string                 `json:"ip_address,omitempty"`
	UserAgent  string                 `json:"user_agent,omitempty"`
	Latency    string                 `json:"latency,omitempty"`
	Body       map[string]interface{} `json:"body"`
}

func customLogMiddleware() gin.HandlerFunc {
	formatter := &httpRequestLog{}

	return func(c *gin.Context) {
		reqTimestamp := time.Now()
		formatter.Body = getReqBody(c)

		c.Next()

		reqLatency := time.Since(reqTimestamp)

		formatter.RequestID = requestid.Get(c)
		formatter.Method = c.Request.Method
		formatter.Path = c.Request.RequestURI
		formatter.StatusCode = c.Writer.Status()
		formatter.IPAddress = c.ClientIP()
		formatter.Latency = reqLatency.String()
		formatter.Timestamp = reqTimestamp.UnixNano()

		byteFormatter, err := json.Marshal(formatter)
		if err != nil {
			log.Fatalf("failed to init custom log middleware. %s\n", err.Error())
		}

		loggingData := string(byteFormatter)
		log.Println(loggingData)
	}
}

func getReqBody(c *gin.Context) map[string]interface{} {
	var bodyBytes []byte

	var err error

	if c.Request.Body == nil {
		bodyBytes = []byte("{}")
	} else {
		bodyBytes, err = io.ReadAll(c.Request.Body)
		if err != nil {
			bodyBytes = []byte("{}")
		}
	}

	bodyMap := make(map[string]interface{})
	_ = json.Unmarshal(bodyBytes, &bodyMap)

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyMap
}
