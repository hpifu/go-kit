package htl

import (
	"bytes"
	"fmt"
	"strconv"
)

func NewBitSet(capacity int) *BitSet {
	return &BitSet{
		bits:     make([]uint64, (capacity-1)/64+1),
		capacity: capacity,
	}
}

type BitSet struct {
	bits     []uint64
	capacity int
}

func (bs *BitSet) Cap(i int) int {
	return bs.capacity
}

func (bs *BitSet) Add(i int) {
	if i >= bs.capacity || i < 0 {
		return
	}
	bs.bits[i/64] |= 1 << (uint64(i) % 64)
}

func (bs *BitSet) Del(i int) {
	if i >= bs.capacity || i < 0 {
		return
	}
	bs.bits[i/64] &= ^(1 << (uint64(i) % 64))
}

func (bs BitSet) Has(i int) bool {
	return bs.bits[i/64]&(1<<(uint64(i)%64)) != 0
}

func (bs BitSet) Empty() bool {
	for _, bit := range bs.bits {
		if bit != 0 {
			return false
		}
	}
	return true
}

func (bs BitSet) String() string {
	var buf bytes.Buffer
	for i := 0; i < len(bs.bits)-1; i++ {
		buf.WriteString(fmt.Sprintf("%064b", bs.bits[i]))
	}
	format := fmt.Sprintf("%%0%db", (bs.capacity-1)%64+1)
	buf.WriteString(fmt.Sprintf(format, bs.bits[len(bs.bits)-1]))
	return buf.String()
}

func (bs BitSet) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('"')
	for i := 0; i < len(bs.bits)-1; i++ {
		buf.WriteString(fmt.Sprintf("%064b", bs.bits[i]))
	}
	format := fmt.Sprintf("%%0%db", (bs.capacity-1)%64+1)
	buf.WriteString(fmt.Sprintf(format, bs.bits[len(bs.bits)-1]))
	buf.WriteByte('"')

	return buf.Bytes(), nil
}

func (bs *BitSet) UnmarshalJSON(buf []byte) error {
	var err error
	str := string(buf[1 : len(buf)-1])

	bs.capacity = len(str)
	bs.bits = make([]uint64, 0, (bs.capacity-1)/64+1)

	i := 0
	for ; i+64 < len(str); i += 64 {
		v, err := strconv.ParseUint(str[i:i+64], 2, 64)
		if err != nil {
			return err
		}
		bs.bits = append(bs.bits, v)
	}
	v, err := strconv.ParseUint(str[i:], 2, 64)
	if err != nil {
		return err
	}
	bs.bits = append(bs.bits, v)

	return err
}
