package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/erabxes/golang-react-todo/server/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the .env")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURL(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb")

	collection := client.Database(dbName).Collection(collName)
	fmt.Println("collection instance created")

}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-url-encoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTasks()
	err := json.NewEncoder(w).Encode(payload) // to encode our payload into our responsewriter w

}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-wwww-form-url-encoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST") // send our body of data as a request with POST http verb
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.ToDoList
	json.NewDecoder(r.Body).Decode(&task) // decode request into this struct task
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {

}

func UndoTask(w http.ResponseWriter, r *http.Request) {

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {

}
