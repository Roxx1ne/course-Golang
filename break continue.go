package main

import "fmt"

func main() {
	//break
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("peruilangan ke", i)
	}

	//continue
	for a := 0; a <= 10; a++ {
		if a%2 == 0 { //kondisi genap, jika angka genap maka di skip
			continue
		}
		fmt.Println("perulangan ke", a)
	}
}
