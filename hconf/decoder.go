package hconf

type Decoder interface {
	Decode(buf []byte) (Storage, error)
}
