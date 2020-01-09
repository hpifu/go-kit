package hashx

import (
	"hash"
	"hash/crc64"
	"hash/fnv"

	"github.com/spaolacci/murmur3"
)

type Hasher64 interface {
	Hash64(buf []byte) uint64
	Hash64S(str string) uint64
}

type hasher64 struct {
	h hash.Hash64
}

func (h hasher64) Hash64(buf []byte) uint64 {
	h.h.Reset()
	_, _ = h.h.Write(buf)
	return h.h.Sum64()
}

func (h hasher64) Hash64S(str string) uint64 {
	return h.Hash64([]byte(str))
}

type HashAlgorithm64 int

const (
	FNV64 HashAlgorithm64 = iota
	FNV64A
	CRC64ISO
	CRC64ECMA
	MURMUR64
)

func NewHasher64(algorithm64 HashAlgorithm64) Hasher64 {
	switch algorithm64 {
	case FNV64:
		return hasher64{h: fnv.New64()}
	case FNV64A:
		return hasher64{h: fnv.New64a()}
	case CRC64ISO:
		return hasher64{h: crc64.New(crc64.MakeTable(crc64.ISO))}
	case CRC64ECMA:
		return hasher64{h: crc64.New(crc64.MakeTable(crc64.ECMA))}
	case MURMUR64:
		return hasher64{h: murmur3.New64()}
	}
	return nil
}
