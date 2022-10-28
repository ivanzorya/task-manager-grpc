package routes

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	tasks "server/tasks"
	models "server/models"
)

var validate = validator.New()

var taskCollection *mongo.Collection = OpenCollection(Client, "tasks")

func CreateTask(task *tasks.CreateTaskRequest) bool {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	dbTask := models.DBTask{
		ID: primitive.NewObjectID(),
		Subject: &task.Task.Subject,
		Done: &task.Task.Done, 
	}

	_, insertErr := taskCollection.InsertOne(ctx, dbTask)
	if insertErr != nil {
		fmt.Println(insertErr)
		return false
	}
	defer cancel()

	return true
}

func GetTasks() []*tasks.Task {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	
	var tasksDB []bson.M

	cursor, err := taskCollection.Find(ctx, bson.M{})

	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	if err = cursor.All(ctx, &tasksDB); err != nil {
		fmt.Println(err)
		return nil
	}

	defer cancel()

	responseTasks := []*tasks.Task{}

	for _, task := range tasksDB {
		bsonBytes, _ := bson.Marshal(task)
		respTmp := &tasks.Task{}
		dbTmp := &models.DBTask{}
		bson.Unmarshal(bsonBytes, dbTmp)
		respTmp.Id = dbTmp.ID.Hex()
		respTmp.Subject = *dbTmp.Subject
		respTmp.Done = *dbTmp.Done
		responseTasks = append(responseTasks, respTmp)
		
	}
	return responseTasks
}

func UpdateTask(task *tasks.UpdateTaskRequest) bool {

	taskID := task.Task.Id
	docID, _ := primitive.ObjectIDFromHex(taskID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	_, err := taskCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"subject":  &task.Task.Subject,
			"done": &task.Task.Done,
		},
	)

	if err != nil {
		fmt.Println(err)
		return false
	}

	defer cancel()

	return true
}

func DeleteTask(task *tasks.DeleteTaskRequest) bool {
	
	taskID := task.Id
	docID, _ := primitive.ObjectIDFromHex(taskID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	_, err := taskCollection.DeleteOne(ctx, bson.M{"_id": docID})
	
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer cancel()

	return true

}
