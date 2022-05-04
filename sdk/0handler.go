// Generated ZEROPS sdk

package sdk

import (
	"net/http"

	"github.com/zeropsio/zerops-go/sdkBase"
)

type Handler struct {
	config      sdkBase.Config
	environment sdkBase.Environment
}

func New(
	config sdkBase.Config,
	httpClient *http.Client,
) Handler {
	return newSdk(
		config,
		sdkBase.NewEnvironment(config, httpClient),
	)
}

func newSdk(
	config sdkBase.Config,
	environment sdkBase.Environment,
) Handler {
	return Handler{
		config:      config,
		environment: environment,
	}
}

func AuthorizeSdk(h Handler, token string) Handler {
	return newSdk(h.config, h.environment.Authorize(token))
}
