package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"todoApp/lib"
)

type Todo struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Text   string `json:"text",omitempty`
	Status bool `json:"status",omitempty`
}

type TodoList[] *Todo

var client *mongo.Client
var collection *mongo.Collection

func init() {
	client = lib.GetClient()
	collection = client.Database("todos").Collection("todolist")
}

func UpdateStatus(id *primitive.ObjectID, status *bool) (Todo, error) {
	findCondition := bson.D{{"_id", &id}}
	updateValue := bson.D{
		{"$set", bson.D{
			{"status", &status},
		}},
	}
	after := options.After
	findOptions := options.FindOneAndUpdateOptions{
		ReturnDocument:&after,
	}
	result := collection.FindOneAndUpdate(context.TODO(), findCondition, updateValue, &findOptions)
	var t Todo
	err := result.Decode(&t)
	if err != nil {
		fmt.Println(err)
		return Todo{}, err
	}
	return t, nil
}

func DeleteTodo(id *primitive.ObjectID) (bool, error) {
	findCondition := bson.D{{"_id", &id}}
	result, err := collection.DeleteOne(context.TODO(), findCondition)
	if err != nil  {
		log.Fatal(err)
		return false, err
	}
	if result.DeletedCount == 0  {
		fmt.Println("No documents found")
		return false, nil
	}
	return true, nil
}

func ListTodo() (TodoList, error) {
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	var results TodoList
	if err != nil {
		log.Fatal(err)
		return TodoList{}, err
	}
	
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var t Todo
		err := cur.Decode(&t)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &t)
	}
	return results,nil
}

func CreateTodo(t *Todo) (bool, error) {
	t.Id = primitive.NewObjectID()
	insertResult, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		fmt.Println("Error in inserting doc", err)
		return false, err
	}
	fmt.Println("successfully inserted doc", insertResult)
	return true, nil
}