package conn

import (
	"net/http"
	"time"

	"github.com/yituoshiniao/kit/xhttp/defaultclient"
)

// HttpClient GO默认 http.Client
type HttpClient struct {
	*http.Client
}

// NewHttpClient // https://developer.apple.com/documentation/appstoreserverapi 接口首页
func NewHttpClient() *HttpClient {
	return &HttpClient{
		Client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: defaultclient.New(
				defaultclient.WithServiceName("httpClient"),
			),
		},
	}
}
