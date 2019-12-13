package henv

import (
	"os"
	"strings"
)

func NewHEnv(prefix string) *HEnv {
	kvs := map[string]string{}
	for _, kv := range os.Environ() {
		idx := strings.Index(kv, "=")
		if idx < 0 {
			continue
		}
		key := kv[:idx]
		val := kv[idx+1:]

		if prefix != "" {
			if strings.HasPrefix(key, prefix+"_") {
				kvs[key[len(prefix)+1:]] = val
			}
		} else {
			kvs[key] = val
		}
	}

	return &HEnv{
		kvs:    kvs,
		prefix: prefix,
	}
}

type HEnv struct {
	prefix string
	kvs    map[string]string
}
