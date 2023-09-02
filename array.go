package main

import "fmt"

func main() {
	//array manual
	var name [3]string

	name[0] = " Naufal"
	name[1] = " Aulio"
	name[2] = " Sopian"

	name[1] = "panjul"

	fmt.Println(name[1])

	//array langsung

	var angka = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(angka)
	fmt.Println(len(angka)) //print panjajng array

}
