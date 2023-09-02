package main

import "fmt"

//defer
func loggging() {
	fmt.Println("selesai memanggil function")
}

func runapplication(value int) {
	defer loggging() // ketika runapplication selesai di jalankan, maka di akhiri dengan run func logging
	fmt.Println("run aplication")
	hasil := 10 / value
	fmt.Println(hasil)
}

func main() {
	runapplication(1)
}
