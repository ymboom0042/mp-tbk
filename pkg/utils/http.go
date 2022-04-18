/**
 * @Author: YMBoom
 * @Description:
 * @File:  http
 * @Version: 1.0.0
 * @Date: 2022/01/13 16:43
 */
package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var HttpRequest *HttpReq

type HttpReq struct {
	Client *http.Client
	mu     sync.Mutex
}

func NewHttpReq(timeout time.Duration) {
	HttpRequest = &HttpReq{
		Client: &http.Client{
			Timeout: timeout * time.Second,
		},
		mu: sync.Mutex{},
	}
}

// Get请求
func (h *HttpReq) Get(url string) (res []byte, err error) {
	h.mu.Lock()
	resp, err := h.Client.Get(url)
	h.mu.Unlock()
	if err != nil {
		return
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Post请求
func (h *HttpReq) Post(reqUrl string, args map[string]string) (res []byte, err error) {
	h.mu.Lock()
	resp, err := h.Client.PostForm(reqUrl, ArgsEncode(args))
	h.mu.Unlock()
	if err != nil {
		return
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func ArgsEncode(args map[string]string) url.Values {
	v := url.Values{}
	for k, arg := range args {
		v.Add(k, arg)
	}

	return v
}
