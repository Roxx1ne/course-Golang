package main

import "fmt"

type Address struct {
	city, province, country string
}

func ChangeCountryToIndonesia(negara *Address) {
	negara.country = "Indonesia"
}

func main() {
	var alamat = Address{
		city:     "bogor",
		province: "Jawa Barat",
		country:  " ",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
