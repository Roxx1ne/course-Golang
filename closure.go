package main

import "fmt"

func main() {
	name := "nopal"

	namalengkap := func() {
		// name = "aulio" //mereplace vairable name di atas
		name := "sopian"  //buat variable baru menggunakan := dan berdiri sendiri tanpa bergantung variabale di atas
		fmt.Println(name) //hanya bisa di akses dalam scope yang sama

		// a := "B"
	}
	// fmt.Println(a) // ini error karena fmt hanya bisa diakses di scope yang sama
	namalengkap()
	fmt.Println(name)
}
