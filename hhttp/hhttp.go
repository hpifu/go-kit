package hhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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

func (h *HttpClient) Do(method string, uri string, header map[string]string, params map[string]interface{}, req interface{}) *HttpResult {
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

	if header != nil {
		for k, v := range header {
			hreq.Header.Add(k, v)
		}
	}

	if params != nil {
		values := &url.Values{}
		for k, v := range params {
			switch v.(type) {
			case string:
				values.Add(k, v.(string))
			case []string:
				for _, i := range v.([]string) {
					values.Add(k, i)
				}
			case int:
				values.Add(k, strconv.Itoa(v.(int)))
			case []int:
				for _, i := range v.([]int) {
					values.Add(k, strconv.Itoa(i))
				}
			default:
				values.Add(k, fmt.Sprintf("%v", v))
			}
		}
		hreq.URL.RawQuery = values.Encode()
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

func (h *HttpClient) GET(uri string, header map[string]string, params map[string]interface{}, req interface{}) *HttpResult {
	return h.Do("GET", uri, header, params, req)
}

func (h *HttpClient) POST(uri string, header map[string]string, params map[string]interface{}, req interface{}) *HttpResult {
	return h.Do("POST", uri, header, params, req)
}
