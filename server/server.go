package server

import (
	"fmt"
	"net/http"
	"urlshort/mapper"
)

type serverContext struct {
	m mapper.Mapper
}

func RunServer(m mapper.Mapper) {
	context := new(serverContext)
	context.m = m

	http.HandleFunc("/register", context.registerHandler)
	http.HandleFunc("/s/", context.redirectHandler)

	fmt.Println("Server running at http://127.0.0.1:1664")
	http.ListenAndServe("127.0.0.1:1664", nil)
}
