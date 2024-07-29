package main

import (
	"fmt"
)

func main() {
	var n int
	var name string
	//take basic input
	fmt.Print("Enter Number your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter Number of subjects: ")
	fmt.Scan(&n)
	var subjects = make(map[string]float32, n)
	fmt.Print("for each subject Enter subject name and Grade separated by space: ")
	for i := 0; i < n; i++ {
		var name string
		var grade float32
		fmt.Scan(&name, &grade)
		if grade < 0 || grade > 100 {
			fmt.Println("Grade should be between 0 and 100")
			i-- //the i variable stays the same
			continue
		}
		subjects[name] = grade
	}
	average := calculateAverage(subjects)

	// print students name and the subjects and their grades and last their average
	fmt.Printf("Student Name: %s\n", name)
	fmt.Println("Subjects and Grades:")
	for subject, grade := range subjects {
		fmt.Printf("%s: %.2f\n", subject, grade)
	}
	fmt.Printf("Average: %.2f\n", average)

}
func calculateAverage(subjects map[string]float32) float32 {
	total := 0.0
	if len(subjects) < 1 {
		return 0.0
	}
	for _, grade := range subjects {
		total += float64(grade)
	}
	return float32(total / float64(len(subjects)))
}
