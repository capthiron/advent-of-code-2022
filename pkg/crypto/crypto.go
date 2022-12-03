package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func HashKey(key string) string {
	hash := md5.Sum([]byte(key))
	return hex.EncodeToString(hash[:])
}
