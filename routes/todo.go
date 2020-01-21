package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"todoApp/models"
)

type updateReqBody struct {
	Id *primitive.ObjectID `json:"_id" bson:"_id"`
	Status *bool `json:"status,omitempty"`
}

type deleteReqBody struct {
	Id *primitive.ObjectID `json:"_id" bson:"_id"`
}

type deleteResBody struct {
	Success bool `json:"success",omitempty`
}

func listAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := models.ListTodo()
	if err != nil {
		fmt.Println("Error in listing docs", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func createTodo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	_, err := models.CreateTodo(&todo)
	
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var u updateReqBody
	_ = json.NewDecoder(r.Body).Decode(&u)
	updatedTodo, err := models.UpdateStatus(u.Id, u.Status)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var d *deleteReqBody
	_ = json.NewDecoder(r.Body).Decode(&d)
	updated, err := models.DeleteTodo(d.Id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := deleteResBody {
		Success : updated,
	}
	json.NewEncoder(w).Encode(res)
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/todos", listAllTodos).Methods("GET")
	r.HandleFunc("/api/create", createTodo).Methods("POST")
	r.HandleFunc("/api/update", updateTodo).Methods("POST")
	r.HandleFunc("/api/delete", deleteTodo).Methods("POST")
	return r
}