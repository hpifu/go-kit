package hashx

import (
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"

	"github.com/spaolacci/murmur3"
)

type Hasher32 interface {
	Hash32(buf []byte) uint32
	Hash32S(str string) uint32
}

type hasher32 struct {
	h hash.Hash32
}

func (h hasher32) Hash32(buf []byte) uint32 {
	h.h.Reset()
	_, _ = h.h.Write(buf)
	return h.h.Sum32()
}

func (h hasher32) Hash32S(str string) uint32 {
	return h.Hash32([]byte(str))
}

type HashAlgorithm32 int

const (
	FNV32 HashAlgorithm32 = iota
	FNV32A
	CRC32IEEE
	CRC32CASTAGNOLI
	CRC32KOOPMAN
	ADLER32
	MURMUR32
)

func NewHasher32(algorithm32 HashAlgorithm32) Hasher32 {
	switch algorithm32 {
	case FNV32:
		return hasher32{h: fnv.New32()}
	case FNV32A:
		return hasher32{h: fnv.New32a()}
	case CRC32IEEE:
		return hasher32{h: crc32.NewIEEE()}
	case CRC32CASTAGNOLI:
		return hasher32{h: crc32.New(crc32.MakeTable(crc32.Castagnoli))}
	case CRC32KOOPMAN:
		return hasher32{h: crc32.New(crc32.MakeTable(crc32.Koopman))}
	case ADLER32:
		return hasher32{h: adler32.New()}
	case MURMUR32:
		return hasher32{h: murmur3.New32()}
	}
	return nil
}
