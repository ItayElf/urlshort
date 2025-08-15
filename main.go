package main

import "fmt"

func main() {
	m := MakeMapper()
	fmt.Println(m)
	m.RegisterRoute("https://gobyexample.com/errors")
	fmt.Println(m)
}
