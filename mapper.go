package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const URL_HASH_LENGTH int = 6

type Mapper struct {
	urls map[string]string
}

func MakeMapper() Mapper {
	return Mapper{make(map[string]string)}
}

func (m *Mapper) RegisterRoute(endpoint string) (string, error) {
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
	return hex.EncodeToString(hash[:])[:URL_HASH_LENGTH]
}
