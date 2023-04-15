package graphql

import (
	"surasithit/gin-graphql-server/graph"
	"surasithit/gin-graphql-server/players"
	"surasithit/gin-graphql-server/teams"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

var resolver *graph.Resolver

func Initialize(playerService players.Service, teamService teams.Service) {
	resolver = &graph.Resolver{
		PlayerService: playerService,
		TeamService:   teamService,
	}
}

// Defining the Graphql handler
func GraphqlHandler(resolver *graph.Resolver) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func PlaygroundHandler(prefixPath string) gin.HandlerFunc {
	h := playground.Handler("GraphQL", prefixPath+"/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
