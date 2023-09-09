package main

import (
	"fmt"
	"reflect"
)

type sample struct {
	Nama string `required:"true" max:"10"`
	age  int    `required:"true" max:"10"`
}

type hewan struct {
	NamaHewan  string `required:"true" max:"10"`
	JenisHewan string `required:"true" max:"10"`
	Habitat    string `required:"true" max:"10"`
}

// validasi library
func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" { //harus ada tag nya di struct biar tervalidasi
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == " " {
				return false
			}

		}
	}
	return true
}

func main() {
	result := sample{"nopal", 21} //inisialisasi struct sample
	sampleType := reflect.TypeOf(result)

	fmt.Println(sampleType.NumField())    //jmlh field di struct
	fmt.Println(sampleType.Field(1).Name) //ambil index dari struct
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	// fmt.Println(sampleType.Field(0).Tag.Get("min")) (print string kosong)

	//cek validasi
	result.Nama = " " //jika string kosong maka false
	fmt.Println(IsValid(result))

	animal := hewan{" ", "mamalia", "darat"} /* jika salah satu ada string kosong maka hasilnya false karena ada tag*/
	fmt.Println(IsValid(animal))

}
