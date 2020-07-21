package docker

import (
	"net/http"

	"github.com/docker/docker/client"

	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
)

var DClient *client.Client
var DClientError error

func Init(host, version string, httpClient *http.Client, httpHeaders map[string]string) {
	DClient, DClientError = client.NewClient(host, version, httpClient, httpHeaders)
	if DClientError != nil {
		xray.ErrMini(DClientError)
	}
}
