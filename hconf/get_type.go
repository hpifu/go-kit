package hconf

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

func (h HConf) GetDefaultInt(keys string, defaultValue ...int) int {
	v, err := h.GetInt(keys)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}

	return v
}

func (h HConf) GetDefaultFloat64(keys string, defaultValue ...float64) float64 {
	v, err := h.GetFloat64(keys)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0.0
		}
		return defaultValue[0]
	}

	return v
}

func (h HConf) GetDefaultString(keys string, defaultValue ...string) string {
	v, err := h.GetString(keys)
	if err != nil {
		if len(defaultValue) == 0 {
			return ""
		}
		return defaultValue[0]
	}

	return v
}

func (h HConf) GetDefaultDuration(keys string, defaultValue ...time.Duration) time.Duration {
	v, err := h.GetDuration(keys)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}

	return v
}

func (h HConf) GetInt(keys string) (int, error) {
	v, err := h.Get(keys)
	if err != nil {
		return 0, err
	}
	return cast.ToIntE(v)
}

func (h HConf) GetFloat64(keys string) (float64, error) {
	v, err := h.Get(keys)
	if err != nil {
		return 0.0, err
	}
	return cast.ToFloat64E(v)
}

func (h HConf) GetString(keys string) (string, error) {
	v, err := h.Get(keys)
	if err != nil {
		return "", err
	}
	return cast.ToStringE(v)
}

func (h HConf) GetDuration(keys string) (time.Duration, error) {
	v, err := h.Get(keys)
	if err != nil {
		return 0, err
	}

	switch v.(type) {
	case string:
		return time.ParseDuration(v.(string))
	case int:
		return time.Duration(v.(int)) * time.Second, nil
	}

	return 0, fmt.Errorf("convert to duration failed")
}
