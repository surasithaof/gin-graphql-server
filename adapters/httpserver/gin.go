package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var server *http.Server
var config Config

type Config struct {
	Mode string `envconfig:"HTTP_MODE" default:"debug"`

	Port   int    `envconfig:"HTTP_PORT" default:"3000"`
	Prefix string `envconfig:"HTTP_PATH_PREFIX" default:""`

	AllowedOrigins []string      `envconfig:"HTTP_ALLOWED_ORIGINS" default:"*"`
	AllowedHeaders []string      `envconfig:"HTTP_ALLOWED_HEADERS" default:"*"`
	CORSMaxAge     time.Duration `envconfig:"HTTP_CORS_MAX_AGE" default:"24h"`
}

func InitGin(_config Config) *gin.Engine {
	config = _config

	if config.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router = gin.New()

	router.Use(
		gin.Logger(),
		gin.Recovery(),
		CORSMiddleware(config),
	)
	return router
}

func Run() *http.Server {
	if router == nil {
		panic("router not initialized")
	}
	server = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: router,
	}
	fmt.Printf("Server is running on port: %s\n", server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("ListenAndServe: %s\n", err)
		}
	}()

	return server
}

func Shutdown(ctx context.Context) {
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %s\n", err)
	}

	fmt.Println("HTTP server has shut down")
}
