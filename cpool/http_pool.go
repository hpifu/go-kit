package connpool

import (
	"net"
	"net/http"
	"time"
)

type HttpPool struct {
	connTimeout time.Duration
	recvTimeout time.Duration
	connQueue   chan *http.Client
}

func NewHttpPool(maxConn int, connTimeout time.Duration, recvTimeout time.Duration) *HttpPool {
	return &HttpPool{
		connTimeout: connTimeout,
		recvTimeout: recvTimeout,
		connQueue:   make(chan *http.Client, maxConn),
	}
}

func (rp *HttpPool) Get() *http.Client {
	select {
	case conn := <-rp.connQueue:
		return conn
	default:
		return &http.Client{
			Transport: &http.Transport{
				Dial: func(netw, addr string) (net.Conn, error) {
					c, err := net.DialTimeout(netw, addr, rp.connTimeout)
					if err != nil {
						return nil, err
					}
					return c, nil
				},
			},
			Timeout: rp.connTimeout + rp.recvTimeout,
		}
	}
}

func (rp *HttpPool) Put(conn *http.Client) {
	select {
	case rp.connQueue <- conn:
		return
	default:
		return
	}
}
