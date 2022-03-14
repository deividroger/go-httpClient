package examples

import (
	"time"

	"github.com/deividroger/go-httpClient/gohttp"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(10 * time.Second).
		SetResponseTimeout(10 * time.Second).
		Build()
	return client
}
