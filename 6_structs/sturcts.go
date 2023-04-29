package main

import (
	"log"
	"time"
)


type User struct{

	// If a variable of function starts with capital letter then it can be accessed outside of the file. It is like public in Java.
	 
	FirstName string
	LastName string
	Age int
	StudentID int
	DateOfBirth time.Time

}

func main() {

	user:= User{
		
	FirstName: "Sakibur",
	LastName: "Rahman",
	}

	log.Println(user.FirstName, user.LastName)

	student:= User{
		FirstName: "Sakib",
		LastName: "Espak",
		Age: 25,
		StudentID: 18101024,
		DateOfBirth: time.Date(1999, time.Month(10), 10, 0, 0, 0, 0, time.UTC),
	}

	log.Println("First Name: ", student.FirstName , "Last Name: ", student.LastName, "Age: ",student.Age, "Student ID", student.StudentID, "Date of Birth", student.DateOfBirth)
}