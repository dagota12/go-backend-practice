package main

import (
	"goPractice/task3/services"
)

func main() {
	// start the library service
	library_service := services.NewLibraryService()
	library_service.Run()

}
