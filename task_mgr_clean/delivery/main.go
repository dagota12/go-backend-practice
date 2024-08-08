package main

import (
	"fmt"
	"goPractice/task_manager/infrastructure"
)

func main() {
	infrastructure.NewMongoClient()
	fmt.Println("Hello world")
}
