package main

import (
	"flag"
	"fmt"
)

func main() {

	host := flag.String("host", "localhost", "silahkan masukan database host")
	user := flag.String("user", "root", "silahkan masukan database user")
	password := flag.String("password", "root", "silahkan masukan database password")
	number := flag.Int("number", 100, "silahkan masukan number dengan benar")

	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*user)
	fmt.Println(*password)
	fmt.Println(*number)
}
