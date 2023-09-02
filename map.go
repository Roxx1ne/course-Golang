package main

import "fmt"

func main() {
	biodata := map[string]string{
		"nama": "naufal ",
		"umur": "21",
		"univ": "esa unggul",
	}
	//menambahkan key "alamat"
	biodata["alamat"] = "jalan mangga no 1/32"

	//delete key umur
	// delete(biodata, "umur")

	//add key
	biodata["kelamin"] = "laki"

	fmt.Println(len(biodata))
	fmt.Println(biodata["nama"])
	fmt.Println(biodata["kelamin"])
	fmt.Println(biodata["umur"])
	fmt.Println(biodata["univ"])
	fmt.Println(biodata["alamat"])
	fmt.Print("\n")

	//nilai ujian
	nilaiakhir := map[string]int{ //string buat tipe data key nya , int buat value nya
		"UAS":   80,
		"UTS":   85,
		"Tugas": 75,
	}
	fmt.Println(len(nilaiakhir))
	fmt.Println(nilaiakhir["UAS"])
	fmt.Println(nilaiakhir["UTS"])
	fmt.Println(nilaiakhir["Tugas"])

	for key, value := range nilaiakhir {
		fmt.Println(key, "=", value) // menampilkan semua data di key dan value
	}
}
