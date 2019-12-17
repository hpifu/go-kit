package hconf

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"sync"
)

func New(decoderType string, providerType string, params ...interface{}) (*HConf, error) {
	provider, err := NewProvider(providerType, params...)
	if err != nil {
		return nil, err
	}
	buf, err := provider.Get()
	if err != nil {
		return nil, err
	}
	decoder, err := NewDecoder(decoderType)
	if err != nil {
		return nil, err
	}
	storage, err := decoder.Decode(buf)
	if err != nil {
		return nil, err
	}
	return &HConf{
		provider: provider,
		storage:  storage,
		decoder:  decoder,
		log:      logrus.New(),
	}, nil
}

type HConf struct {
	provider Provider
	storage  Storage
	decoder  Decoder

	envPrefix string
	handlers  []OnChangeHandler
	log       *logrus.Logger

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func (h *HConf) SetLogger(log *logrus.Logger) {
	h.log = log
}

const MapMod = 1
const ArrMod = 2

type KeyInfo struct {
	key string
	idx int
	mod int
}

func parseKey(keys string, separator string) ([]*KeyInfo, error) {
	var infos []*KeyInfo
	for _, key := range strings.Split(keys, separator) {
		fields := strings.Split(key, "[")
		if len(fields[0]) != 0 {
			infos = append(infos, &KeyInfo{
				key: fields[0],
				mod: MapMod,
			})
		}
		for i := 1; i < len(fields); i++ {
			if !strings.HasSuffix(fields[i], "]") {
				return nil, fmt.Errorf("invalid format. key: [%v]", key)
			}
			field := fields[i][0 : len(fields[i])-1]
			if len(field) == 0 {
				return nil, fmt.Errorf("index should not be empty. key: [%v]", key)
			}
			idx, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("index should be a number. key: [%v]", key)
			}
			infos = append(infos, &KeyInfo{
				idx: idx,
				mod: ArrMod,
			})
		}
	}

	return infos, nil
}

func (h HConf) Get(key string) (interface{}, error) {
	return h.storage.Get(key)
}

func (h *HConf) Set(key string, val interface{}) error {
	return h.storage.Set(key, val)
}

func (h HConf) Unmarshal(v interface{}) error {
	return h.storage.Unmarshal(v)
}

func (h HConf) Sub(key string) (*HConf, error) {
	s, err := h.storage.Sub(key)
	if err != nil {
		return nil, err
	}

	return &HConf{
		storage: s,
	}, nil
}

func (h *HConf) SetEnvPrefix(prefix string) {
	h.envPrefix = prefix
}

func (h *HConf) BindEnv(key string, env ...string) error {
	envkey := key
	if len(env) != 0 {
		envkey = env[0]
	} else {
		envkey = strings.ReplaceAll(envkey, ".", "_")
		envkey = strings.ReplaceAll(envkey, "[", "_")
		envkey = strings.ReplaceAll(envkey, "]", "")
		envkey = strings.ToUpper(envkey)
		if h.envPrefix != "" {
			envkey = h.envPrefix + "_" + envkey
		}
	}
	val := os.Getenv(envkey)

	if val == "" {
		return nil
	}

	v, err := h.Get(key)
	if err != nil {
		return h.Set(key, val)
	}

	switch v.(type) {
	case string:
		return h.Set(key, val)
	case float64:
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		return h.Set(key, f)
	case int:
		i, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		return h.Set(key, i)
	case float32:
		f, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return err
		}
		return h.Set(key, float32(f))
	}

	return fmt.Errorf("val type can not bind env, val: [%v]", v)
}
