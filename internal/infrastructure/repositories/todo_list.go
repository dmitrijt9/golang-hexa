package repositories

import (
	"context"
	"hexa-example-go/internal/domain/adapters"
	"hexa-example-go/internal/domain/entities"
	"log"
	"math/rand"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type TodoListRepo struct {
	logger   zap.Logger
	mongoCli mongo.Client
}

func NewTodoListRepo(logger zap.Logger, mongoCli mongo.Client) *TodoListRepo {
	return &TodoListRepo{
		logger:   logger,
		mongoCli: mongoCli,
	}
}

func (r *TodoListRepo) Save(toSave adapters.TodoListToSave) (*entities.TodoList, error) {
	dbName := "test"
	collName := "test"
	coll := r.mongoCli.Database(dbName).Collection(collName)

	//context will time out after 30s
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	insert := bson.M{"_id": toSave.Name + "_" + strconv.Itoa(rand.Intn(1000)), "value": toSave}
	res, err := coll.InsertOne(ctx, insert)

	if err != nil {
		log.Panic(err)
	}

	return &entities.TodoList{
		Id:   res.InsertedID.(int),
		Name: toSave.Name,
	}, nil
}

// TODO: implement
func (r *TodoListRepo) GetByName(name string) (*entities.TodoList, error) {
	panic("not implemented yet")
}

func (r *TodoListRepo) List() ([]*entities.TodoList, error) {
	panic("not implemented yet")
}
