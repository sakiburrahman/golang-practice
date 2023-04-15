package main

import "log"

func main() {
	var itemName string
	itemName = "Desktop Computer"
	log.Println("The initial item is", itemName)
	changeItem(&itemName)
	log.Println("The new item is", itemName)
}

func changeItem(item *string){
	log.Println("The value of item is", item)
	newItem:= "Laptop"
	*item = newItem
}