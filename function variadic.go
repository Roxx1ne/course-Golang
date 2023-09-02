package main

import "fmt"

func sumALL(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	hasil := sumALL(90, 10) // func sumALL di masukan ke variable hasil
	fmt.Println(hasil)

	//memasukan slice ke variadic
	slice := []int{10, 20, 30, 40, 50}
	hasil = sumALL(slice...)
	fmt.Println(hasil)
}
