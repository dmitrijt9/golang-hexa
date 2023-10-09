package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"hexa-example-go/internal/app/domain/entities"
	"hexa-example-go/internal/app/domain/out_ports"
	"hexa-example-go/internal/logger"
)

type TodoRepo struct {
	logger   logger.Logger
	mongoCli mongo.Client
}

func (t TodoRepo) Save(toSave out_ports.TodoToSave) (*entities.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoRepo) List() ([]*entities.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func NewTodoRepo(logger logger.Logger, mongoCli mongo.Client) out_ports.TodoRepository {
	return &TodoRepo{
		logger:   logger,
		mongoCli: mongoCli,
	}
}
