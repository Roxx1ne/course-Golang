package main

import "fmt"

//person
type name interface {
	getName() string
	getUmur() int
}

func sayName(people name) {
	fmt.Println("ini", people.getName())
	fmt.Println("berumur", people.getUmur())
}

type orang struct {
	Nama string
	umur int
}

func (biodata orang) getName() string {
	return biodata.Nama
}

func (biodata orang) getUmur() int {
	return biodata.umur
}

//hewan
type peliharaan interface {
	getJenis() string
	getNama() string
	getJumlahKaki() int
}

func sayanimal(hewan peliharaan) {
	fmt.Println("jenis hewan : ", hewan.getJenis())
	fmt.Println("nama hewan  : ", hewan.getNama())
	fmt.Println("jumlah kaki : ", hewan.getJumlahKaki())
}

type animal struct {
	Jenis      string
	Nama       string
	JumlahKaki int
}

func (hewan animal) getNama() string {
	return hewan.Nama
}
func (hewan animal) getJenis() string {
	return hewan.Jenis
}
func (hewan animal) getJumlahKaki() int {
	return hewan.JumlahKaki
}

func main() {
	var budi orang
	budi.Nama = "budi"
	budi.umur = 20

	sayName(budi)

	hewan := animal{
		Nama:       "caty",
		Jenis:      "kucing",
		JumlahKaki: 4,
	}

	sayanimal(hewan)
}
