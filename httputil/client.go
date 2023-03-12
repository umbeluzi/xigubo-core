package httputil

import (
	"crypto/tls"
	"net/http"
	"time"
)

var (
	defaultTimeout = 10 * time.Second
)

func NewClient(opts ...ClientOption) (*http.Client, error) {
	client := &http.Client{
		Timeout: defaultTimeout,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

type ClientOption func(*http.Client)

func WithCert(cert []byte) ClientOption {
	return func(c *http.Client) {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{},
		}
	}
}

func WithCACert(cert []byte) ClientOption {
	return func(c *http.Client) {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{},
		}
	}
}
