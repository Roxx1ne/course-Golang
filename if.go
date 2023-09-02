package main

import "fmt"

func main() {
	name := "NAUFAL"

	if name == "naufal" {
		fmt.Println("data benar")
	} else {
		fmt.Println("data salah")
	}

	if length := len(name); length == 6 {
		fmt.Println("panjang data sudah sesuai")
	} else {
		fmt.Println("data terlalu panjang")
	}

	nilaiuas := 80

	if nilaiuas < 80 {
		fmt.Println("mohon maaf anda tidak lulus")
	} else if nilaiuas == 80 {
		fmt.Println("selamat anda lulus")
	} else {
		fmt.Println("mohon maaf anda tidak lulus")
	}
}
