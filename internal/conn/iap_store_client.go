package conn

import (
	"context"

	"github.com/awa/go-iap/appstore/api"
)

type IapStoreClient struct {
	httpClient *HttpClient
}

func NewIapStoreClient(
	httpClient *HttpClient,
) *IapStoreClient {
	return &IapStoreClient{
		httpClient: httpClient,
	}
}

func (i *IapStoreClient) StoreClientWithHTTPClient(ctx context.Context, conf *api.StoreConfig) *api.StoreClient {
	return api.NewStoreClientWithHTTPClient(conf, i.httpClient.Client)
}
