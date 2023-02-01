package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type TodoRepo struct {
	logger   zap.Logger
	mongoCli mongo.Client
}

func NewTodoRepo(logger zap.Logger, mongoCli mongo.Client) *TodoRepo {
	return &TodoRepo{
		logger:   logger,
		mongoCli: mongoCli,
	}
}

// TODO: implement
// func (r *TodoRepo) Save(toSave adapters.TodoListToSave) (*entities.TodoList, error) {
// 	dbName := "test"
// 	collName := "test"
// 	coll := r.mongoCli.Database(dbName).Collection(collName)

// 	//context will time out after 30s
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

// 	defer cancel()

// 	insert := bson.M{"_id": toSave.Name + "_" + strconv.Itoa(rand.Intn(1000)), "value": toSave}
// 	res, err := coll.InsertOne(ctx, insert)

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	return &entities.TodoList{
// 		Id:   res.InsertedID.(int),
// 		Name: toSave.Name,
// 	}, nil
// }

// func (r *TodoRepo) GetByName(name string) (*entities.TodoList, error)
// func (r *TodoRepo) List() ([]*entities.TodoList, error)
