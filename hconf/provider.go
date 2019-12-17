package hconf

import (
	"context"
	"fmt"
	"reflect"
)

func NewProvider(name string, params ...interface{}) (Provider, error) {
	switch name {
	case "local":
		if len(params) != 1 {
			return nil, fmt.Errorf("params should equal 1 for %v provider, got %v", name, len(params))
		}
		filename, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("expect a string parameter for %v provider, got [%v]", name, reflect.TypeOf(params[0]))
		}
		return NewLocalProvider(filename)
	case "consul":
		if len(params) != 2 {
			return nil, fmt.Errorf("params should equal 2 for %v provider, got %v", name, len(params))
		}
		address, ok := params[1].(string)
		if !ok {
			return nil, fmt.Errorf("expect a string parameter for %v provider address, got %v", name, params[0])
		}
		key, ok := params[1].(string)
		if !ok {
			return nil, fmt.Errorf("expect a string parameter for %v provider key, got %v", name, params[1])
		}
		return NewConsulProvider(address, key)
	}

	return nil, fmt.Errorf("unsupport provider %v", name)
}

type Provider interface {
	Events() <-chan struct{}
	Errors() <-chan error
	Get() ([]byte, error)
	EventLoop(ctx context.Context)
}
