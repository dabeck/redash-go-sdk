package administration

import (
	"github.com/dabeck/redash-go-sdk/gen/client"
	admin "github.com/dabeck/redash-go-sdk/gen/client/administration"
)

type RequestWrapper struct {
	httpClient *client.Redashclient
	host       string
	opts       []admin.ClientOption
}

func NewRequestWrapper(httpClient *client.Redashclient, opts ...admin.ClientOption) *RequestWrapper {
	return &RequestWrapper{httpClient: httpClient, opts: opts}
}

// Ping verify connection initialized
func (requestWrapper *RequestWrapper) Ping() (err error) {
	params := admin.NewGetPingParams()
	_, err = requestWrapper.httpClient.Administration.GetPing(params, nil, requestWrapper.opts...)
	return err
}
