package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"hexa-example-go/internal/app/domain/entities"
	"hexa-example-go/internal/app/domain/out_ports"
)

type TodoRepo struct {
	logger   zap.Logger
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

func NewTodoRepo(logger zap.Logger, mongoCli mongo.Client) out_ports.TodoRepository {
	return &TodoRepo{
		logger:   logger,
		mongoCli: mongoCli,
	}
}
