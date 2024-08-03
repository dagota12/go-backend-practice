package main

import (
	"goPractice/task_manager/config"
	"goPractice/task_manager/controllers"
	"goPractice/task_manager/data"
	"goPractice/task_manager/router"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Title       string    `bson:"title,omitempty" json:"title"`
	Description string    `bson:"description,omitempty" json:"description"`
	DueDate     time.Time `bson:"due_date,omitempty" json:"due_date"`
	Status      string    `bson:"status,omitempty" json:"status"`
}

var tasks = map[string]Task{
	"1": {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	"2": {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	"3": {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

type Trainer struct {
	Name string `bson:"name,omitempty" json:"name"`
	Age  int    `bson:"age,omitempty" json:"age"`
	City string `bson:"city,omitempty" json:"city"`
}
type Server struct {
	Router *gin.Engine
	DB     *mongo.Database
	Port   string
}

func (s *Server) Run() {
	log.Println("Starting server on port: 8080")
	s.Router = gin.Default()
	s.DB = config.NewMongoClient().Database("test")

	task_service := data.NewTaskService(s.DB.Collection("tasks"))
	// // task_service.Tasks = tasks
	task_controller := controllers.NewTaskController(task_service)
	router.SetUpRouter(s.Router, task_controller)
	s.Router.Run(":8080")

}
func main() {

	server := Server{Port: "8080"}
	server.Run()
	// router_inst := gin.Default()
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// client := config.NewMongoClient()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db := client.Database("test")
	// trainees := db.Collection("trainees")
	// fmt.Println(trainees)

	// ash := Trainer{"Ash", 10, "Pallet Town"}
	// misty := Trainer{"Misty", 10, "Cerulean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}

	//FILTER
	// filter := bson.M{"name": "Ash"}
	// update := bson.D{
	// 	{Key: "$inc", Value: bson.D{
	// 		{Key: "age", Value: 1},
	// 	}},
	// }
	// upd, err := trainees.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Matched %v documents and updated %v documents.\n", upd.MatchedCount, upd.ModifiedCount)

	//FIND
	// var trainer Trainer

	// err := trainees.FindOne(context.TODO(), filter).Decode(&trainer)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Found a single document: %+v\n", trainer)

	// // ENV Variable loading
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found")
	// }
	// fmt.Println("DBURI:", os.Getenv("DB_URI"))

	//INSERT MANY
	// trainers := []interface{}{misty, brock}
	// result, err := trainees.InsertMany(context.TODO(), trainers)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result.InsertedIDs...)

	//INSERT ONE
	// result, err := trainees.InsertOne(context.TODO(), ash)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted a single document: ", result.InsertedID)

	// task_service := data.NewTaskService()
	// // task_service.Tasks = tasks
	// task_controller := controllers.NewTaskController(task_service)
	// router.SetUpRouter(router_inst, task_controller)
	// router_inst.Run(":8080")
}
