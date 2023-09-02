package main

import "fmt"

func main() {
	type alamat string // kata alamat jadi alias dari tipedata string

	var alamatpengguna alamat = " jalan mangga blok 1 no 26"
	fmt.Println(alamatpengguna)
}
