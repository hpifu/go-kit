package hashx

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Hash(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func BKDRHash(str string) uint64 {
	seed := uint64(131) // 31 131 1313 13131 131313 etc..
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint64(str[i])
	}
	return hash & 0x7FFFFFFF
}

func DJBHash(str string) uint64  {
	hash := uint64(5381)
	for i := 0; i < len(str); i++ {
		hash += (hash << 5) + uint64(str[i])
	}
	return hash & 0x7FFFFFFF
}

func ELFHash(str string) uint64  {
	x := uint64(0)
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = (hash << 4) + uint64(str[i])
		if (x & hash & 0xF0000000) != 0 {
			hash ^= x >> 24
			hash &= ^x
		}
	}
	return hash & 0x7FFFFFFF
}

func FNVHash(str string) uint64  {
	fnvPrime := uint64(0x811C9DC5)
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash *= fnvPrime
		hash ^= uint64(str[i])
	}
	return hash & 0x7FFFFFFF
}

