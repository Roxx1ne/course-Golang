package main

import "fmt"

func ups() interface{} {
	return "ups"
}

func contoh(contoh int) interface{} {
	if contoh == 1 {
		return 1
	} else if contoh == 2 {
		return true
	} else {
		return "contoh interface kosong"
	}
}

func main() {
	kosong := ups()
	fmt.Println(kosong)

	interfacekosong := contoh(3)
	/* atau sama aja
	var interfacekosong interface{} = contoh(3) */
	fmt.Println(interfacekosong)

}
