package server

import (
	"errors"
	"net/http"
	"strings"
)

func (context *serverContext) redirectHandler(w http.ResponseWriter, r *http.Request) {
	hash, err := getHash(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	route, err := context.m.GetRouteByHash(hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, route, http.StatusSeeOther)
}

func getHash(url string) (string, error) {
	url = strings.TrimPrefix(url, "/s/")
	url_parts := strings.Split(url, "/")
	if len(url_parts) != 1 {
		return "", errors.New("url must be in format '/s/<hash>'")
	}
	return url_parts[0], nil
}
