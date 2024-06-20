package src

type ICityUsecases interface {
}

type CityUsecases struct {
}

func NewCityUsecases() *CityUsecases {
	return &CityUsecases{}
}
