package main

import "fmt"

//type declaration
type Block = func(string) bool

func register(name string, blacklist Block) {
	if blacklist(name) {
		fmt.Println("youre blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	register("admin", blacklist)
	register("nopal", blacklist)
}
