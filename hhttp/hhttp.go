package hhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hpifu/go-kit/cpool"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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
	param := &url.Values{}

	val := reflect.ValueOf(req).Elem()
	for i := 0; i < val.Type().NumField(); i++ {
		key := strings.Split(val.Type().Field(i).Tag.Get("json"), ",")[0]
		switch val.Type().Field(i).Tag.Get("http") {
		case "body":
			body[key] = val.Field(i).String()
		case "param":
			param.Add(key, val.Field(i).String())
		}
	}

	buf, _ := json.Marshal(body)

	hreq, err := http.NewRequest(method, uri, bytes.NewReader(buf))
	if err != nil {
		return err
	}

	hreq.URL.RawQuery = param.Encode()

	hres, err := client.Do(hreq)
	if err != nil {
		return err
	}
	buf, err = ioutil.ReadAll(hres.Body)
	if err != nil {
		return err
	}
	defer hres.Body.Close()

	if err := json.Unmarshal(buf, res); err != nil {
		return fmt.Errorf("res: [%v], err: [%v]", string(buf), err)
	}

	h.pool.Put(client)

	return nil
}

func (h *HttpClient) GET(uri string, req interface{}, res interface{}) error {
	return h.Do("GET", uri, req, res)
}

func (h *HttpClient) POST(uri string, req interface{}, res interface{}) error {
	return h.Do("POST", uri, req, res)
}
