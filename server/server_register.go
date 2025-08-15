package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type registerRouteParameters struct {
	Endpoint string `json:"endpoint"`
}

func (context *serverContext) registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parameters := &registerRouteParameters{}
	err := json.NewDecoder(r.Body).Decode(parameters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := context.m.RegisterRoute(parameters.Endpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("http://127.0.0.1:1664/s/%s", hash)
	log.Printf("Registered endpoint at %s", url)
	fmt.Fprint(w, url)
}
