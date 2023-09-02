package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := pembagian(100, 0)

	if err == nil {
		fmt.Println("hasil = ", hasil)
	} else {
		fmt.Println("errors = ", err.Error())
	}
}
