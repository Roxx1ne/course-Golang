package main

import "fmt"

//recover
func endapp() {
	message := recover()
	if message == nil {
		fmt.Println("error dengan message: ", message)
	}
	fmt.Println("aplikasi selesai di jalankan")

}

func runapp(errors bool) {
	defer endapp()
	if errors {
		panic("APLIKASI ERROR")
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	runapp(false)
	fmt.Println("test") // jika menggunakan recover, ketika terjadi panic, script di bawwah runapp masih bisa dijalankan
}
