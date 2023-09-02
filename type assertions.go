package main

import "fmt"

func contoh() interface{} {
	return 12
}

func main() {
	// result := contoh()
	// resultString := result.(int)
	// // resultString := result.(string) // ini akan error karena tidak sesuai tipe data
	// fmt.Println(resultString)

	//menggunakan switch
	hasil := contoh()
	switch value := hasil.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("integer", value)
	case bool:
		fmt.Println("bool", value)
	default:
		fmt.Println("tipe data tidak terbaca")
	}

}
