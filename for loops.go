package main

import "fmt"

func main() {
	angka := 20

	for angka >= 1 {
		fmt.Println("perulangan ke", angka)
		angka--
	}

	//versi singkat
	for angka := 20; angka >= 1; angka-- {
		fmt.Println("perulangan ke", angka)
	}

	// for range
	biodata := []string{"nopal", "aulio", "panjul"}

	for i := 0; i < len(biodata); i++ {
		fmt.Println(biodata[i]) // mengambil semua data di index
	}
	for i, value := range biodata { //for i bisa di ganti for[_] menandakan variable tidak dipake
		fmt.Println("index", i, "=", value)
	}

}
