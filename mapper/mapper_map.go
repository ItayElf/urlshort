package mapper

import (
	"errors"
)

type MapMapper struct {
	urls map[string]string
}

func (m *MapMapper) RegisterRoute(endpoint string) (string, error) {
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

func NewMapMapper() *MapMapper {
	m := new(MapMapper)
	m.urls = make(map[string]string)
	return m
}
