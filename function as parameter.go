package main

import "fmt"

//buat alias
type Filter func(string) string

func sayhellowithfilter(name string, filter Filter) {
	namefiltered := filter(name)
	fmt.Println("hello", namefiltered)
}

//pemrosessan filter
func spamfiltered(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}

}

func main() {
	sayhellowithfilter("anjing", spamfiltered) //jika "anjing" maka hello ...
	sayhellowithfilter("nopal", spamfiltered)  //spamfilter akan memfilter setiap string di println
}
