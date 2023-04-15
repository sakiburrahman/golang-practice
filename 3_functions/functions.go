package main

import "fmt"

func main() {

	name, id, cGPA := studentInfo()
	fmt.Println("The student name is", name)
	fmt.Println("Student ID is", id)
	fmt.Println("cGPA is", cGPA)
	

}

func studentInfo() (string, int, float64){
	return "Espak", 0510124, 3.97
}