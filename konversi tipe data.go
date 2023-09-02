package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai64)

	var name = "panjjul"
	var a = name[2]
	var aString = string(a)
	fmt.Println(aString)
}
