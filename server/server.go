package server

import (
	"fmt"
	"net/http"
	"urlshort/mapper"
)

func RunServer(m mapper.Mapper) {
	fmt.Println("Server running at http://127.0.0.1:1664")
	http.ListenAndServe("127.0.0.1:1664", nil)
}
