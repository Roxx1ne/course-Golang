package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User //alias

// func untuk package sort
func (value UserSlice) Len() int {
	return len(value) //untuk print panjang struct
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age //Compare dari yang terkecil(BEBAS)
}
func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"nopal", 21},
		{"budi", 25},
		{"tono", 19},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
