package main

import (
	"fmt"
	"container/List"
)
func main(){
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)
	intList.PushFront(45)

	for element := intList.Front(); element!=nil; element= element.Next() {
		fmt.Println(element.Value.(int))
	}
	
}