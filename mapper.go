package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const urlHashLength = 6

type Mapper interface {
	registerRoute(endpoint string) (string, error)
}

type MapMapper struct {
	urls map[string]string
}

func MakeMapMapper() MapMapper {
	return MapMapper{make(map[string]string)}
}

func (m *MapMapper) registerRoute(endpoint string) (string, error) {
	hash := generateEndpointHash(endpoint)

	if v, exists := m.urls[hash]; exists {
		if v == endpoint {
			return hash, nil
		}
		return "", errors.New("route is already taken")
	}

	m.urls[hash] = endpoint

	return hash, nil
}

func generateEndpointHash(endpoint string) string {
	hash := md5.Sum([]byte(endpoint))
	return hex.EncodeToString(hash[:])[:urlHashLength]
}
