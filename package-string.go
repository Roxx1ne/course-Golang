package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim("     Naufal Aulio    ", " "))                         //hps spasi->(bebas)
	fmt.Println(strings.ToLower("NAUFAL AULIO SOPIAN"))                             //lowercase
	fmt.Println(strings.ToUpper("naufal aulio sopian"))                             //uppercase
	fmt.Println(strings.Contains("naufal aulio", "aulio"))                          //cek apakah ada string (aulio)
	fmt.Println(strings.ReplaceAll("Naufal Naufal Naufal Naufal", "Naufal", "lio")) //Naufal replace (lio)
	fmt.Println(strings.Split("Naufal Aulio Sopian", "a"))                          //hapus kata (a)
}
