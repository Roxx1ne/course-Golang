package main

import "fmt"

func factorial(value int) int {
	if value == 1 { //kondisi jika nilai ==1 maka loop stop, di 1
		return 1
	} else {
		return value * factorial(value-1)
	}
}

func main() {
	hasil := factorial(6)
	fmt.Println(hasil)

}
