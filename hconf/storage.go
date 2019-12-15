package hconf

type Storage interface {
	Get(key string) (interface{}, error)
	Set(key string, val interface{}) error
	Unmarshal(v interface{}) error
	Sub(key string) (Storage, error)
}
