package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("arguments : ")
	fmt.Println(args)
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])

	//os.hostname
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("hostname : ", name)
	} else {
		fmt.Println("error ", err.Error())
	}

	//os.GetEnv
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	//run di terminal bash
	fmt.Println(username)
	fmt.Println(password)

}
