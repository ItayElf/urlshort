package mapper

import (
	"crypto/md5"
	"encoding/hex"
)

const urlHashLength = 6

type Mapper interface {
	RegisterRoute(endpoint string) (string, error)
	GetRouteByHash(hash string) (string, error)
}

func generateEndpointHash(endpoint string) string {
	hash := md5.Sum([]byte(endpoint))
	return hex.EncodeToString(hash[:])[:urlHashLength]
}
