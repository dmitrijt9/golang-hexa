package container

import (
	"hexa-example-go/internal/app/domain/in_ports"
	"hexa-example-go/internal/app/domain/services"
	"hexa-example-go/internal/app/infrastructure/clients/mongo"
	"hexa-example-go/internal/app/infrastructure/repositories"
	"hexa-example-go/internal/config"
	"hexa-example-go/internal/logger"
)

func New() Container {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger := logger.New(logger.Config{
		Level:  logger.Level(conf.Logger.Level),
		Format: logger.Format(conf.Logger.Format),
	})

	mongoClient := mongo.Connect(conf.Mongo)
	defer mongo.Disconnect(&mongoClient)

	todoListRepo := repositories.NewTodoListRepo(logger, mongoClient)
	todoListService := services.NewTodoListService(logger, todoListRepo)

	return Container{
		Config:          conf,
		Logger:          logger,
		TodoListService: todoListService,
	}

}

type Container struct {
	Config config.HexaExampleConfig
	Logger logger.Logger

	TodoListService in_ports.TodoListService
}
