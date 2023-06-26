package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/surasithaof/gin-graphql-server/adapters/db"
	"github.com/surasithaof/gin-graphql-server/adapters/graphql"
	"github.com/surasithaof/gin-graphql-server/adapters/httpserver"
	"github.com/surasithaof/gin-graphql-server/graph"
	"github.com/surasithaof/gin-graphql-server/graph/loader"
	"github.com/surasithaof/gin-graphql-server/players"
	"github.com/surasithaof/gin-graphql-server/teams"

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

	rGroup.GET("/health", healthCheck)

	database := db.Connect(config.Database)
	playerStore := players.Initialize(database)
	teamStore := teams.Initialize(database)

	// Initialize loader
	ldrs := loader.NewLoaders(playerStore, teamStore)

	gqlResolver := graph.Initialize(playerStore, teamStore)

	rGroup.POST("/query", loader.Middleware(ldrs), graphql.GraphqlHandler(gqlResolver))
	rGroup.GET("/", graphql.PlaygroundHandler(config.HttpServer.Prefix))
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
