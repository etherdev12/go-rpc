package rpc

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"runtime"
)

// RpcOption is a functional option.
type RpcOption func(*RpcConfig)

type RpcConfig struct {
	Headers map[string]string
}

func NewConfig(opts ...RpcOption) RpcConfig {
	c := RpcConfig{
		Headers: make(map[string]string),
	}

	c.Headers["Rpc-Agent"] = runtime.GOOS

	for _, o := range opts {
		o(&c)
	}

	return c
}

// WithCustomHeaders sets headers of http request.
func WithCustomHeaders(key, value string) RpcOption {
	return func(c *RpcConfig) {
		c.Headers[key] = value
	}
}

func RpcQuery(URL string, opts ...RpcOption) ([]byte, error) {
	r, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	config := NewConfig(opts...)
	for k, v := range config.Headers {
		r.Header.Add(k, v)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	return body, err
}
