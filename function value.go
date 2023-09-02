package main

import "fmt"

func getGoodBye(name string) string {
	return "hello" + name
}

func main() {
	goodbye := getGoodBye //func getGoddBye di masukan ke variable goodbye
	fmt.Println(goodbye(" nopal "))
}
