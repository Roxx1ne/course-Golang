package main

import "fmt"

func newmap(name string) map[string]string { //return map, key string, value string
	if name == "" {
		return nil //nil sama seperti null
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	hasil := newmap("")
	fmt.Println(hasil)

	//cek data
	var name map[string]string = newmap("")

	if name == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(name)
	}
}
