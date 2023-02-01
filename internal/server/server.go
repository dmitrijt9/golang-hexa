package server

import (
	graph "hexa-example-go/internal/app/interface/graphql"
	c "hexa-example-go/internal/container"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func New(container c.Container) *gin.Engine {
	r := gin.Default()
	r.POST("/api", graphqlHandler(container))
	r.GET("/", playgroundHandler())

	// init staic server for Voyager GQL explorer
	r.StaticFile("/explorer", "./explorer/voyager.html")

	return r
}

// Defining the Graphql handler
func graphqlHandler(deps c.Container) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			TodoListService: deps.TodoListService,
		}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the GQL Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/api")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
