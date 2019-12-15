package hconf

import (
	"context"
	"github.com/hashicorp/consul/api"
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
		for {
			kv, meta, err := p.client.KV().Get(p.key, &api.QueryOptions{WaitIndex: p.lastIndex, WaitHash: p.lastHash})
			if err != nil {
				p.errors <- err
				continue
			}

			p.lastIndex = meta.LastIndex
			if p.lastHash == meta.LastContentHash {
				continue
			}

			p.lastHash = meta.LastContentHash
			p.buf = kv.Value
			p.events <- struct{}{}
		}
	}()
}
