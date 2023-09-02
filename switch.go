package main

import "fmt"

func main() {
	name := "Nopal"

	switch name {
	case "Nopal":
		fmt.Println("halo nopal")
	case "Budi":
		fmt.Println("halo budi")
	default:
		fmt.Println("tidak ada nama yang cocok")
	}

	//short statement
	switch panjangkata := len(name); panjangkata >= 5 {
	case true:
		fmt.Println("panjang kata sudah tepat")
	case false:
		fmt.Println("panjanga kata terlalu panjang")
	}

	//switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("panjang kata terlalu panjang")
	case length >= 5:
		fmt.Println("panjang kata sudah sesuai ")
	default:
		fmt.Println("panjang kata kurang dari 5")
	}
}
