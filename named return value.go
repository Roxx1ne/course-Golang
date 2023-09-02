package main

import "fmt"

func namedvalues() (nama string, umur int, alamat string, universitas string) {
	nama = "Naufal Aulio"
	umur = 21
	alamat = "jalan condet no 21 "
	universitas = "esa unggul "
	return
}

func main() {
	a, b, c, d := namedvalues()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
