package main

import "fmt"

//panic
func endapp() {
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
	runapp(true)
}
