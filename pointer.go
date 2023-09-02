package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	//pass by value
	alamat := Address{"Tangerang", "Kab.Tangerang", "indonesia"}
	alamat2 := alamat

	alamat2.province = "tangerang selatan"

	fmt.Println(alamat)
	fmt.Println(alamat2)
	fmt.Println("\n")

	//pas by reference
	address1 := Address{"Kemang", "Jakarta Selatan", "Indonesia"}
	address2 := &address1 //address 2 == address 1 atau (pointer)
	address3 := &address1 //buat test jika ada data baru yang pointer ke address 1

	//ganti 1 field
	address2.city = "lebak bulus"

	//ganti semua data
	// address2 = &Address{"Malang", "Jawa timur", "Indonesia"}

	fmt.Println(address1)
	// fmt.Println(address2) // data akan sama dengan address 1 karena , bersifat pointer
	fmt.Println("\n")

	//operator *
	*address2 = Address{"bandung", "Jawa barat", "Indonesia"} //address 1 mengikuti data di address2

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3) // address 3 akan sama datanya , karena address 1 pake operator *

	var address4 *Address = new(Address) /*variable address4 itu mengambil pointer dari struct address,
	tapi membuat address baru yang isi datanya kosong*/
	// address4.country = "singapore"
	fmt.Println(address4)
}
