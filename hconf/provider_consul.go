package hconf

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"time"
)

type ConsulProvider struct {
	key       string
	client    *api.Client
	lastIndex uint64
	lastHash  string

	events chan struct{}
	errors chan error
	buf    []byte
}

func NewConsulProvider(address string, key string) (*ConsulProvider, error) {
	config := api.DefaultConfig()
	config.Address = address
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	kv, meta, err := client.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}
	if kv == nil {
		return nil, fmt.Errorf("config [%v] not found", key)
	}

	return &ConsulProvider{
		client:    client,
		key:       key,
		lastIndex: meta.LastIndex,
		lastHash:  meta.LastContentHash,
		buf:       kv.Value,
		events:    make(chan struct{}, 10),
		errors:    make(chan error, 10),
	}, nil
}

func (p *ConsulProvider) Events() <-chan struct{} {
	return p.events
}

func (p *ConsulProvider) Errors() <-chan error {
	return p.errors
}

func (p *ConsulProvider) Get() ([]byte, error) {
	return p.buf, nil
}

func (p *ConsulProvider) EventLoop(ctx context.Context) {
	go func() {
		go func() {
			sleepTimeOnError := 200 * time.Millisecond
			for {
				kv, meta, err := p.client.KV().Get(p.key, &api.QueryOptions{WaitIndex: p.lastIndex, WaitHash: p.lastHash})
				if kv == nil {
					err = fmt.Errorf("config [%v] not found", p.key)
				}
				if err != nil {
					p.errors <- err
					time.Sleep(sleepTimeOnError)
					sleepTimeOnError *= 2
					if sleepTimeOnError > 60*time.Second {
						sleepTimeOnError = 60 * time.Second
					}
					continue
				}

				p.events <- struct{}{}
				sleepTimeOnError = 200 * time.Millisecond
				p.lastIndex = meta.LastIndex
				p.lastHash = meta.LastContentHash
				p.buf = kv.Value
			}
		}()
	out:
		for {
			select {
			case <-ctx.Done():
				break out
			}
		}
	}()
}
