package hconf

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"github.com/prometheus/common/log"
	"path/filepath"
	"time"
)

func (h *HConf) Stop() {
	h.cancel()
	h.wg.Wait()
}

func (h *HConf) Watch() error {
	h.ctx, h.cancel = context.WithCancel(context.Background())

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	abs, err := filepath.Abs(h.filename)
	if err != nil {
		return err
	}
	if err := watcher.Add(filepath.Dir(abs)); err != nil {
		return err
	}

	go func() {
		h.wg.Add(1)
	out:
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Rename) == 0 {
					continue
				}
				if event.Name != abs {
					continue
				}
				for len(watcher.Events) != 0 {
					<-watcher.Events
				}
				nh, err := NewHConfWithFile(h.filename)
				if err != nil {
					log.Warnf("reload file failed. filename: [%v]", h.filename)
					continue
				}
				log.Infof("reload file success. filename: [%v]", h.filename)
				h.data = nh.data
				for _, handler := range h.handlers {
					handler(h)
				}
			case err := <-watcher.Errors:
				log.Warnf("watcher error [%v]", err)
			case <-time.Tick(time.Duration(8) * time.Second):
				log.Info("tick")
			case <-h.ctx.Done():
				log.Info("stop watch. exit")
				break out
			}
		}
		h.wg.Done()
	}()

	return nil
}

type OnChangeHandler func(h *HConf)

func (h *HConf) RegisterOnChangeHandler(handler OnChangeHandler) {
	h.handlers = append(h.handlers, handler)
}
