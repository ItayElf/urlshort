package main

import (
	"fmt"
	"urlshort/mapper"
)

func main() {
	var m mapper.Mapper = mapper.NewMapMapper()
	fmt.Println(m)
	m.RegisterRoute("https://gobyexample.com/errors")
	fmt.Println(m)
}
