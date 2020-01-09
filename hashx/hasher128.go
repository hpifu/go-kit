package hashx

import (
	"encoding/binary"
	"hash"
	"hash/fnv"

	"github.com/spaolacci/murmur3"
)

type Hasher128 interface {
	Hash128(buf []byte) (uint64, uint64)
	Hash128S(str string) (uint64, uint64)
}

type murmur3Hasher128 struct{}

func (murmur3Hasher128) Hash128(buf []byte) (uint64, uint64) {
	return murmur3.Sum128(buf)
}

func (murmur3Hasher128) Hash128S(str string) (uint64, uint64) {
	return murmur3.Sum128([]byte(str))
}

type fnvHasher128 struct {
	h hash.Hash
}

func (h fnvHasher128) Hash128(buf []byte) (uint64, uint64) {
	h.h.Reset()
	_, _ = h.h.Write(buf)
	b := h.h.Sum(nil)
	return binary.BigEndian.Uint64(b[0:8]), binary.BigEndian.Uint64(b[8:16])
}

func (h fnvHasher128) Hash128S(str string) (uint64, uint64) {
	return h.Hash128([]byte(str))
}

type HashAlgorithm128 int

const (
	FNV128 HashAlgorithm128 = iota
	FNV128A
	MURMUR128
)

func NewHasher128(algorithm128 HashAlgorithm128) Hasher128 {
	switch algorithm128 {
	case FNV128:
		return fnvHasher128{h: fnv.New128()}
	case FNV128A:
		return fnvHasher128{h: fnv.New128a()}
	case MURMUR128:
		return murmur3Hasher128{}
	}
	return nil
}
