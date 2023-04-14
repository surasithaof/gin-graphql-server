package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"surasithit/gin-graphql-server/adapters/httpserver"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Start() error {
	_config, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	router := httpserver.InitGin(_config.HttpServer)

	httpserver.Run()

	initialApp(router, _config)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	httpserver.Shutdown(ctx)

	return nil
}

func initialApp(router *gin.Engine, config *Config) {
	rGroup := router.Group(config.HttpServer.Prefix)

	rGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
