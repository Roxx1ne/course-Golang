package main

import "fmt"

type CustomerID struct {
	Name, Address, Gender string
	Age                   int
}

//struct method
func (pelanggan CustomerID) sayHi(name string) {
	fmt.Println("hello", name, "my name is", pelanggan.Name)

}

func (sapa CustomerID) sayMabar() {
	fmt.Println("p mabar", sapa.Name)
}

func main() {
	var nopal CustomerID

	nopal.Name = "Naufal Aulio"
	nopal.Address = "jalan merak"
	nopal.Gender = " Laki - laki"
	nopal.Age = 20

	fmt.Println(nopal)

	//struct method
	nopal.sayHi("ahmad")
	nopal.sayMabar()

	//struct literals
	panjul := CustomerID{
		Name:    "Panjul Riski",
		Address: "jalan bangau",
		Gender:  "laki - laki",
		Age:     18,
	}
	fmt.Println(panjul)
}
