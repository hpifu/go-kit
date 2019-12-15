package hconf

import (
	"context"
	"github.com/prometheus/common/log"
	"time"
)

func (h *HConf) Stop() {
	h.cancel()
	h.wg.Wait()
}

func (h *HConf) Watch() error {
	h.ctx, h.cancel = context.WithCancel(context.Background())
	h.provider.EventLoop(h.ctx)

	go func() {
		h.wg.Add(1)
	out:
		for {
			select {
			case <-h.provider.Events():
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
			case err := <-h.provider.Errors():
				log.Warnf("provider error [%v]", err)
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
