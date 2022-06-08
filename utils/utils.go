package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func Min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
func Max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
