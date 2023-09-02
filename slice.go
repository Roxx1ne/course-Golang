package main

import "fmt"

func main() {
	var bulan = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice = bulan[2:5] //ambil data dari array 2 dan 5 sebagai akhir
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice)) // hitung capacity dari array 2

	slice[0] = "lio" // 0 di slide adalah 2 pada [2:5]
	fmt.Println(slice)

	slice2 := bulan[2:4]
	fmt.Println(slice2)

	slice3 := append(slice2, "nopal")
	slice3[1] = "Test" //mengubah indeks 0 yaitu november dari slice2

	fmt.Println(slice3)
	fmt.Println(bulan)

	//cara simple
	slicebaru := make([]string, 2, 5)
	slicebaru[0] = "budi"
	slicebaru[1] = "santoso"

	fmt.Println(slicebaru)
	fmt.Println(len(slicebaru))
	fmt.Println(cap(slicebaru))

	//copy slice
	copyslice := make([]string, len(slicebaru), cap(slicebaru))
	copy(copyslice, slicebaru)
	fmt.Println(copyslice)

	//perbedaan array dan slice
	//[... / 5 ] itu array , kalo slice itu []
	iniarray := [...]int{1, 2, 3, 4, 5}
	inislice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniarray)
	fmt.Println(inislice)

}
