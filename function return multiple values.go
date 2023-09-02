package main

import "fmt"

func biodata2() (string, string, int) {
	return "Naufal Aulio", "jalan apel", 21
}

func main() {
	name, addres, _ := biodata2() // _ digunakan untuk igonre tipe data int diatas
	fmt.Println(name, addres)
}
