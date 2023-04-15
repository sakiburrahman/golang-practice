package main

import "log"

// Global variable 
var colorName = "Green"


func main() {

	// Local variable inside the function
	var newColorName = "Red"
	log.Println("New color name is", newColorName)

	// Function is printing the parameter value. value overriding
	changeColor(" Yellow")

	// Function is printing the global colorName value.
	changeToOrangeColor("find")
}

func changeColor(colorName string) (string){

	log.Println("The color name from changeColor function is", colorName)
	return colorName
}

func changeToOrangeColor(colorNameOr string) (string){

	log.Println("The color name from changeToOrangeColor function is", colorName)
	return colorNameOr
}