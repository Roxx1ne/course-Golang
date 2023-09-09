package main

import (
	"fmt"
	"strconv"
)

func main() {
	bool, err := strconv.ParseBool("true") //mengubah string ke bool
	if err == nil {
		fmt.Println(bool)
	} else {
		fmt.Println(err.Error())
	}

	//biary(2), desimal(10), oktal(8), hexa(16)
	number, err := strconv.ParseInt("1000000", 10, 64) //10 (baseint), 64(size)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 2) //tidak ada kondisi error karena sudah fix.
	fmt.Println(value)
}
