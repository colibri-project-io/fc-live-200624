package src

import (
	"net/http"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
)

type CityController struct {
	usecase ICityUsecases
}

func NewCityController() *CityController {
	return &CityController{
		usecase: NewCityUsecases(),
	}
}

func (c *CityController) Routes() []restserver.Route {
	return []restserver.Route{
		{
			URI:      "cities",
			Method:   http.MethodGet,
			Function: c.GetAll,
			Prefix:   restserver.PublicApi,
		},
		{
			URI:      "cities/{id}",
			Method:   http.MethodGet,
			Function: c.GetById,
			Prefix:   restserver.PublicApi,
		},
		{
			URI:      "cities/cep/{cep}",
			Method:   http.MethodGet,
			Function: c.GetByCep,
			Prefix:   restserver.PublicApi,
		},
		{
			URI:      "cities",
			Method:   http.MethodPost,
			Function: c.CreateCity,
			Prefix:   restserver.PublicApi,
		},
		{
			URI:      "cities/{id}",
			Method:   http.MethodDelete,
			Function: c.DeleteCityById,
			Prefix:   restserver.PublicApi,
		},
	}
}

func (c *CityController) GetAll(wctx restserver.WebContext) {
	wctx.EmptyResponse(http.StatusNotImplemented)
}

func (c *CityController) GetById(wctx restserver.WebContext) {
	wctx.EmptyResponse(http.StatusNotImplemented)
}

func (c *CityController) GetByCep(wctx restserver.WebContext) {
	wctx.EmptyResponse(http.StatusNotImplemented)
}

func (c *CityController) CreateCity(wctx restserver.WebContext) {
	wctx.EmptyResponse(http.StatusNotImplemented)
}

func (c *CityController) DeleteCityById(wctx restserver.WebContext) {
	wctx.EmptyResponse(http.StatusNotImplemented)
}
