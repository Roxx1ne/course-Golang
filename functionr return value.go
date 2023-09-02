package main

import "fmt"

func returnfunction(test string) string {
	if test == "" {
		return "kenalan bro"
	} else {
		return "haloo" + test
	}
}
func main() {
	fmt.Println(returnfunction(" nopal "))

}
