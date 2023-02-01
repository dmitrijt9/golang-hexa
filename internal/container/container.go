package container

import (
	"hexa-example-go/internal/app/domain/services"
	"hexa-example-go/internal/app/infrastructure/repositories"
	"hexa-example-go/internal/config"
	"hexa-example-go/internal/logger"
	"hexa-example-go/internal/mongo"

	"go.uber.org/zap"
)

func New() Container {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger := logger.InitLogger(conf.Logger)

	mongoClient := mongo.Connect(conf.Mongo)
	defer mongo.Disconnect(&mongoClient)

	todoListRepo := repositories.NewTodoListRepo(*logger, mongoClient)
	todoListService := services.NewTodoListService(*logger, todoListRepo)

	return Container{
		Config:          conf,
		Logger:          *logger,
		TodoListService: todoListService,
	}

}

type Container struct {
	Config config.HexaExampleConfig
	Logger zap.Logger

	TodoListService services.TodoListService
}
