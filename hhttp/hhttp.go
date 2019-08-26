package hhttp

import (
	"fmt"
	"github.com/hpifu/go-kit/cpool"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

type HttpClient struct {
	pool *cpool.HttpPool
}

func NewHttpClient(maxConn int, connTimeout time.Duration, recvTimeout time.Duration) *HttpClient {
	return &HttpClient{
		pool: cpool.NewHttpPool(maxConn, connTimeout, recvTimeout),
	}
}

func (h *HttpClient) Do(method string, uri string, req interface{}, res interface{}) error {
	client := h.pool.Get()

	body := map[string]interface{}{}

	val := reflect.ValueOf(req)
	for i := 0; i < val.Type().NumField(); i++ {
		switch val.Type().Field(i).Tag.Get("http") {
		case "body":
			body[val.Type().Field(i).Name] = "1"
		}
	}
	fmt.Println(body)
	hreq, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return err
	}

	q := &url.Values{}
	q.Add("token", "123")
	hreq.URL.RawQuery = q.Encode()

	hres, err := client.Do(hreq)
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadAll(hres.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(buf))
	defer hres.Body.Close()

	h.pool.Put(client)

	return nil
}

func (h *HttpClient) GET(uri string, req interface{}, res interface{}) error {
	return h.Do("GET", uri, req, res)
}

func (h *HttpClient) POST(uri string, req interface{}, res interface{}) error {
	return h.Do("POST", uri, req, res)
}
