package main

import (
	"fmt"
	"strconv"
)

func main() {
	var angka int // satu variable hanya bisa menampung 1 tipe data
	//var angka string = ini tidak bisa karena variabel bersifat unique

	angka = 12 + 40 - 5/2
	fmt.Println(angka)

	name := "naufal aulio sopian" //:= inisialisasi var dan hanya berlaku di awal
	fmt.Println(name)

	name = "saya adalah mahasiswa"
	fmt.Println(name)

	umur := 21
	umurString := strconv.Itoa(umur) // konversi umur jadi string, Itoa = integer to string
	fmt.Println("umur saya adalah " + "" + umurString + " tahun ")

	var boolean = true // secara otomatis mendeteksi tipe data
	fmt.Println(boolean)
}
