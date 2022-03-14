package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers           http.Header
	maxIdleConnection int
	connectionTimeout time.Duration
	responseTimeout   time.Duration
	disableTimeouts   bool
}

type ClientBuilder interface {
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(i int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	SetHeaders(headers http.Header) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
		client:  &http.Client{},
	}

	return &client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}
func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}
func (c *clientBuilder) SetMaxIdleConnections(i int) ClientBuilder {
	c.maxIdleConnection = i
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}
