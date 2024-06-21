package src

import (
	"context"
	"errors"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/base/logging"
	"github.com/google/uuid"
)

type ICityUsecases interface {
	GetAll(ctx context.Context, params *CityPageParams) (CityPageResponse, error)
	GetById(ctx context.Context, id uuid.UUID) (*CityResponse, error)
	GetByCep(ctx context.Context, cep string) (*CityByCepResponse, error)
	Create(ctx context.Context, model *CityCreateRequest) (*CityResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type CityUsecases struct {
	cityRepository ICityRepository
	viaCepGateway  IViaCepGateway
}

func NewCityUsecases() *CityUsecases {
	return &CityUsecases{
		cityRepository: NewCityDBRepository(),
		viaCepGateway:  NewViaCepRestGateway(),
	}
}

func (u *CityUsecases) GetAll(ctx context.Context, params *CityPageParams) (CityPageResponse, error) {
	result, err := u.cityRepository.FindAll(ctx, params)
	if err != nil {
		logging.Error("error in FindAll: %+v", err)
		return nil, errors.New(ErrOnFindAllCity)
	}

	return result, nil
}

func (u *CityUsecases) GetById(ctx context.Context, id uuid.UUID) (*CityResponse, error) {
	exists, err := u.cityRepository.ExistsById(ctx, id)
	if err != nil {
		logging.Error("error in ExistsById: %+v", err)
		return nil, errors.New(ErrOnExistsCityById)
	}

	if !*exists {
		return nil, errors.New(ErrCityIdNotFound)
	}

	result, err := u.cityRepository.FindById(ctx, id)
	if err != nil {
		logging.Error("error in FindById: %+v", err)
		return nil, errors.New(ErrOnFindCityById)
	}

	return result, nil
}
func (u *CityUsecases) GetByCep(ctx context.Context, cep string) (*CityByCepResponse, error) {
	result, err := u.viaCepGateway.FindCityByCep(ctx, cep)
	if err != nil {
		logging.Error("error in FindCityByCep: %+v", err)
		return nil, errors.New(ErrOnFindCityByCep)
	}

	return result.ToCityByResponse(), nil
}
func (u *CityUsecases) Create(ctx context.Context, model *CityCreateRequest) (*CityResponse, error) {
	exists, err := u.cityRepository.ExistsByUniqueKey(ctx, model.Name, model.UF)
	if err != nil {
		logging.Error("error in ExistsByUniqueKey: %+v", err)
		return nil, errors.New(ErrOnExistsCityById)
	}

	if *exists {
		return nil, errors.New(ErrCityAlreadyExists)
	}

	result, err := u.cityRepository.Insert(ctx, model)
	if err != nil {
		logging.Error("error in Insert: %+v", err)
		return nil, errors.New(ErrOnInsertCity)
	}

	return result, nil
}
func (u *CityUsecases) Delete(ctx context.Context, id uuid.UUID) error {
	exists, err := u.cityRepository.ExistsById(ctx, id)
	if err != nil {
		logging.Error("error in ExistsById: %+v", err)
		return errors.New(ErrOnExistsCityById)
	}

	if !*exists {
		return errors.New(ErrCityIdNotFound)
	}

	err = u.cityRepository.Delete(ctx, id)
	if err != nil {
		logging.Error("error in Delete: %+v", err)
		return errors.New(ErrOnDeleteCityById)
	}

	return nil
}
