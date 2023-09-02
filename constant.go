package main

import (
	"fmt"
	"strconv"
)

func main() {
	const (
		alamat = " graha pesona blok w14/26 "
		umur   = 21
		univ   = " esa unggul "
	)

	// alamat = "jalan bangau " (error karena const tidak bisa diubah isinya)

	fmt.Println(alamat + strconv.Itoa(umur) + univ)
}
