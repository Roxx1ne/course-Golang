package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushFront("nopal")
	data.PushBack("aulio")
	data.PushBack("sopian")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
	fmt.Println("\n")

	//iterate
	//dpn ke blkg
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	//blkg ke dpn
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

}
