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

func (m *MapMapper) GetRouteByHash(hash string) (string, error) {
	if v, exists := m.urls[hash]; exists {
		return v, nil
	}
	return "", errors.New("no route was registered for requested hash")
}

func NewMapMapper() *MapMapper {
	m := new(MapMapper)
	m.urls = make(map[string]string)
	return m
}
