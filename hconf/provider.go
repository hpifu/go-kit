package hconf

import "context"

type Provider interface {
	Events() <-chan struct{}
	Errors() <-chan error
	Get() ([]byte, error)
	EventLoop(ctx context.Context)
}
