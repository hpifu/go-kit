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
				for len(h.provider.Events()) != 0 {
					<-h.provider.Events()
				}
				buf, err := h.provider.Get()
				if err != nil {
					log.Warnf("provider get failed. err: [%v]", err)
					continue
				}
				storage, err := h.decoder.Decode(buf)
				if err != nil {
					log.Warnf("decoder decode failed. err: [%v]", err)
				}
				log.Infof("reload config success. storage: %v", h.storage)
				h.storage = storage
				for _, handler := range h.handlers {
					handler(h)
				}
			case err := <-h.provider.Errors():
				log.Warnf("provider error [%v]", err)
			case <-time.Tick(time.Duration(300) * time.Second):
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
