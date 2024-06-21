package src

import (
	"time"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/base/types"
	"github.com/google/uuid"
)

type CityPageResponse *types.Page[CityResponse]

type CityResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	UF        string    `json:"uf"`
	CreatedAt time.Time `json:"createdAt"`
}

type CityPageParams struct {
	Page uint16 `form:"page" validate:"required"`
	Size uint16 `form:"pageSize" validate:"required"`
	Name string `form:"name"`
	UF   string `form:"uf"`
}

type CityCreateRequest struct {
	Name string `json:"name" validate:"required"`
	UF   string `json:"uf" validate:"required"`
}

type ViaCepCityResponse struct {
	Name string `json:"localidade"`
	UF   string `json:"uf"`
}

type CityByCepResponse struct {
	Name string `json:"name"`
	UF   string `json:"uf"`
}

func (model *ViaCepCityResponse) ToCityByResponse() *CityByCepResponse {
	return &CityByCepResponse{
		Name: model.Name,
		UF:   model.UF,
	}
}
