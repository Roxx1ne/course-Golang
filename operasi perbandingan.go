package main

import "fmt"

func main() {
	var string1 = "Nopall"
	var string2 = "Nopall"

	var result = string1 != string2
	fmt.Println(result)

	var UAS = 90
	var ABSEN = 70

	// //ketentuan lulus/tidak lulus
	// var ujian = UAS > 80
	// var presensi = ABSEN > 75

	// var hasil = ujian || presensi
	// fmt.Println(hasil)

	//inisialisasi langsung

	fmt.Println(UAS > 80 || ABSEN > 75)

}
