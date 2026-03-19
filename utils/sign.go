package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Hex(input string) string {
	sum := md5.Sum([]byte(input))
	return hex.EncodeToString(sum[:])
}
