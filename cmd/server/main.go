package main

import (
	"hexa-example-go/internal/domain/services"
	"hexa-example-go/internal/infrastructure/repositories"
	graph "hexa-example-go/internal/interface/graphql"
	"hexa-example-go/internal/pkg/config"
	"hexa-example-go/internal/pkg/logger"
	"hexa-example-go/internal/pkg/mongo"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("Cannot load app config:", err)
	}

	loggerConf := conf.Logger
	logger := logger.InitLogger(logger.LoggerConfig{
		Level:           loggerConf.Level,
		ConsoleEnabled:  loggerConf.ConsoleEnabled,
		FilebeatEnabled: loggerConf.FilebeatEnabled,
		FilebeatUrl:     loggerConf.FilebeatUrl,
		FilebeatIndex:   loggerConf.FilebeatIndex,
		FileBeatAppName: loggerConf.FileBeatAppName,
	})

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	go handleKillSignal(interrupt, *logger)

	mongoClient := mongo.Connect(conf.Mongo)
	defer mongo.Disconnect(&mongoClient)

	todoListRepo := repositories.NewTodoListRepo(*logger, mongoClient)
	todoListService := services.NewTodoListService(*logger, todoListRepo)

	r := initServer(ServerDependencies{
		TodoListService: *todoListService,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	r.Run(host + ":" + port)
}

type ServerDependencies struct {
	TodoListService services.TodoListService
}

func initServer(deps ServerDependencies) *gin.Engine {

	r := gin.Default()
	r.POST("/api", graphqlHandler(deps))
	r.GET("/", playgroundHandler())

	// init staic server for Voyager GQL explorer
	r.StaticFile("/explorer", "./explorer/voyager.html")

	return r
}

func handleKillSignal(interrupt chan os.Signal, logger zap.Logger) {
	for {
		killSignal := <-interrupt
		// TODO: How do we want to handle signals? what should we close, wait for it to end?
		switch killSignal {
		case syscall.SIGINT:
			logger.Info("Got SIGINT...")
			os.Exit(0)
		case syscall.SIGTERM:
			logger.Info("Got SIGTERM...")
			os.Exit(0)
		}
	}
}

// Defining the Graphql handler
func graphqlHandler(deps ServerDependencies) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
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
