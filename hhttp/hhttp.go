package hhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/hpifu/go-kit/cpool"
)

type HttpClient struct {
	pool *cpool.HttpPool
}

type HttpResult struct {
	Status int
	Res    []byte
	Err    error
}

func (hr *HttpResult) String(res *string) error {
	if hr.Err != nil {
		return hr.Err
	}
	*res = string(hr.Res)

	return nil
}

func (hr *HttpResult) Interface(res interface{}) error {
	if hr.Err != nil {
		return hr.Err
	}
	if err := json.Unmarshal(hr.Res, res); err != nil {
		return fmt.Errorf("res: [%v], err: [%v]", string(hr.Res), err)
	}

	return nil
}

func NewHttpClient(maxConn int, connTimeout time.Duration, recvTimeout time.Duration) *HttpClient {
	return &HttpClient{
		pool: cpool.NewHttpPool(maxConn, connTimeout, recvTimeout),
	}
}

func (h *HttpClient) Do(method string, uri string, params map[string]string, req interface{}) *HttpResult {
	client := h.pool.Get()

	var reader io.Reader
	if req != nil {
		buf, _ := json.Marshal(req)
		reader = bytes.NewReader(buf)
	}
	hreq, err := http.NewRequest(method, uri, reader)
	if err != nil {
		return &HttpResult{
			Err: err,
		}
	}

	if params != nil {
		param := &url.Values{}
		for k, v := range params {
			param.Add(k, v)
		}
		hreq.URL.RawQuery = param.Encode()
	}

	hres, err := client.Do(hreq)
	if err != nil {
		return &HttpResult{
			Err: err,
		}
	}
	buf, err := ioutil.ReadAll(hres.Body)
	if err != nil {
		return &HttpResult{
			Err: err,
		}
	}
	defer hres.Body.Close()

	h.pool.Put(client)

	return &HttpResult{
		Status: hres.StatusCode,
		Res:    buf,
	}
}

func (h *HttpClient) GET(uri string, params map[string]string, req interface{}) *HttpResult {
	return h.Do("GET", uri, params, req)
}

func (h *HttpClient) POST(uri string, params map[string]string, req interface{}) *HttpResult {
	return h.Do("POST", uri, params, req)
}
