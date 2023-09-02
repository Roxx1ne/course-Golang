package helper

import "fmt"

//penjelasan package dan access modifier
var Application = "GOLANG"
var version = "2.4.6.1"

func Hello(name string) { //kapital (H) bisa di akses di luar package helper
	fmt.Println("hello", name)
}

func sayGoodBye(name string) { // huruf (s)ay tidak bisa di akses karena huruf kecil
	fmt.Println("goodbye", name)
}
