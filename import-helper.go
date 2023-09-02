package main

import (
	"Golang-dasar/helper"
	"fmt"
)

func main() {
	helper.Hello("nopal")
	// helper.sayGoodBye("nopal") error, karena menggunakan huruf kecil
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) error
}
