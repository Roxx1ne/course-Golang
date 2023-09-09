package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.8))  //cenderung ke atas maka di bulatkan ke atas
	fmt.Println(math.Round(1.3))  //cenderung ke bawah maka di bulatkan ke bawah
	fmt.Println(math.Floor(1.8))  //paksa di bulatkan ke bawah
	fmt.Println(math.Ceil(1.8))   //paksa di bulatkan ke atas
	fmt.Println(math.Max(10, 20)) //print 20
	fmt.Println(math.Min(10, 20)) //print 10
}
