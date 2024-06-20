package src

const (
	viaCepURL      string = "VIACEP_URL"
	viaCepPath     string = "/ws/%s/json/"
	viaCepError    string = "error: %+v"
	cepCachePrefix string = "cep-%v"
)

type IViaCepGateway interface {
}

type ViaCepRestGateway struct {
}

func NewViaCepRestGateway() *ViaCepRestGateway {
	return &ViaCepRestGateway{}
}
