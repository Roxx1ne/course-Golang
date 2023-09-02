package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) person() {
	man.Name = " Mr. " + man.Name
}

func main() {
	nopal := Man{"nopal"}
	nopal.person()

	fmt.Println(nopal.Name)
}
