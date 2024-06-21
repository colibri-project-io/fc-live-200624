package src

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/cacheDB"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restclient"
)

const (
	viaCepURL      string = "VIACEP_URL"
	viaCepPath     string = "/ws/%s/json"
	viaCepError    string = "error: %+v"
	cepCachePrefix string = "cep-%v"
)

type IViaCepGateway interface {
	FindCityByCep(ctx context.Context, cep string) (*ViaCepCityResponse, error)
}

type ViaCepRestGateway struct {
	client *restclient.RestClient
}

func NewViaCepRestGateway() *ViaCepRestGateway {
	return &ViaCepRestGateway{
		client: restclient.NewRestClient(&restclient.RestClientConfig{
			Name:    "via-cep-rest-client",
			BaseURL: os.Getenv(viaCepURL),
			Timeout: 10,
		}),
	}
}

func (g *ViaCepRestGateway) FindCityByCep(ctx context.Context, cep string) (*ViaCepCityResponse, error) {
	response := restclient.Request[ViaCepCityResponse, any]{
		Ctx:        ctx,
		Client:     g.client,
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf(viaCepPath, cep),
		Cache:      cacheDB.NewCache[ViaCepCityResponse](fmt.Sprintf(cepCachePrefix, cep), 10*time.Second),
	}.Call()

	if response.ErrorBody() != nil {
		return nil, fmt.Errorf(viaCepError, response.ErrorBody())
	}

	if response.Error() != nil {
		return nil, response.Error()
	}

	return response.SuccessBody(), nil
}
