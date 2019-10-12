package hrand

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func NewToken() string {
	buf := make([]byte, 32)
	token := make([]byte, 16)
	rand.New(rand.NewSource(time.Now().UnixNano())).Read(token)
	hex.Encode(buf, token)
	return string(buf)
}

func NewAuthCode() string {
	return fmt.Sprintf("%06d", int(rand.NewSource(time.Now().UnixNano()).(rand.Source64).Uint64()%1000000))
}
