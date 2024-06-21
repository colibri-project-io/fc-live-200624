package src

import (
	"net/http"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"github.com/google/uuid"
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
	var params CityPageParams
	if err := wctx.DecodeQueryParams(&params); err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	result, err := c.usecase.GetAll(wctx.Context(), &params)
	if err != nil {
		wctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	wctx.JsonResponse(http.StatusOK, result)
}

func (c *CityController) GetById(wctx restserver.WebContext) {
	paramId, err := uuid.Parse(wctx.PathParam("id"))
	if err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	result, err := c.usecase.GetById(wctx.Context(), paramId)
	if err != nil {
		if err.Error() == ErrCityIdNotFound {
			wctx.ErrorResponse(http.StatusNotFound, err)
			return
		}

		wctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	wctx.JsonResponse(http.StatusOK, result)

}

func (c *CityController) GetByCep(wctx restserver.WebContext) {
	result, err := c.usecase.GetByCep(wctx.Context(), wctx.PathParam("cep"))
	if err != nil {
		wctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	wctx.JsonResponse(http.StatusOK, result)
}

func (c *CityController) CreateCity(wctx restserver.WebContext) {
	var body CityCreateRequest
	if err := wctx.DecodeBody(&body); err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	result, err := c.usecase.Create(wctx.Context(), &body)
	if err != nil {
		if err.Error() == ErrCityAlreadyExists {
			wctx.ErrorResponse(http.StatusConflict, err)
			return
		}

		wctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	wctx.JsonResponse(http.StatusCreated, result)
}

func (c *CityController) DeleteCityById(wctx restserver.WebContext) {
	paramId, err := uuid.Parse(wctx.PathParam("id"))
	if err != nil {
		wctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	if err := c.usecase.Delete(wctx.Context(), paramId); err != nil {
		if err.Error() == ErrCityIdNotFound {
			wctx.ErrorResponse(http.StatusNotFound, err)
			return
		}

		wctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	wctx.EmptyResponse(http.StatusNoContent)
}
