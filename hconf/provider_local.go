package hconf

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"os"
	"path/filepath"
)

type LocalProvider struct {
	abs     string
	watcher *fsnotify.Watcher

	events chan struct{}
	errors chan error
}

func NewLocalProvider(filename string) (*LocalProvider, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	abs, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	if err := watcher.Add(filepath.Dir(abs)); err != nil {
		return nil, err
	}

	return &LocalProvider{
		watcher: watcher,
		abs:     abs,
		events:  make(chan struct{}, 10),
		errors:  make(chan error, 10),
	}, nil
}

func (p *LocalProvider) Events() <-chan struct{} {
	return p.events
}

func (p *LocalProvider) Errors() <-chan error {
	return p.errors
}

func (p *LocalProvider) Get() ([]byte, error) {
	fp, err := os.Open(p.abs)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	return ioutil.ReadAll(fp)
}

func (p *LocalProvider) EventLoop(ctx context.Context) {
	go func() {
	out:
		for {
			select {
			case event := <-p.watcher.Events:
				if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Rename) == 0 {
					continue
				}
				if event.Name != p.abs {
					continue
				}
				for len(p.watcher.Events) != 0 {
					<-p.watcher.Events
				}
				p.events <- struct{}{}
			case err := <-p.watcher.Errors:
				p.errors <- err
			case <-ctx.Done():
				break out
			}
		}
	}()
}
