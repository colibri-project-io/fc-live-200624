package src

import (
	"context"
	"fmt"
	"time"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/base/types"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/cacheDB"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/sqlDB"
	"github.com/google/uuid"
)

const (
	findAllquery string = `
		SELECT id, name, uf, created_at 
		FROM cities
		WHERE 1=1
		AND ($1 = '' OR LOWER(name) LIKE LOWER(CONCAT('%', $1, '%')))
		AND ($2 = '' OR LOWER(uf) LIKE LOWER(CONCAT('%', $2, '%')))`
	findByIdQuery          string = "SELECT id, name, uf, created_at FROM cities WHERE id = $1"
	existsByIdQuery        string = "SELECT EXISTS(SELECT FROM cities WHERE id = $1)"
	existsByUniqueKeyQuery string = "SELECT EXISTS(SELECT FROM cities WHERE name = $1 AND uf = $2)"
	insertQuery            string = "INSERT INTO cities (name, uf) VALUES ($1, $2) RETURNING id, name, uf, created_at"
	deleteQuery            string = "DELETE FROM cities WHERE id = $1"
	cityCachePrefix        string = "city-%v"
)

type ICityRepository interface {
	FindAll(ctx context.Context, params *CityPageParams) (CityPageResponse, error)
	FindById(ctx context.Context, id uuid.UUID) (*CityResponse, error)
	ExistsById(ctx context.Context, id uuid.UUID) (*bool, error)
	ExistsByUniqueKey(ctx context.Context, name string, uf string) (*bool, error)
	Insert(ctx context.Context, model *CityCreateRequest) (*CityResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type CityDBRepository struct{}

func NewCityDBRepository() *CityDBRepository {
	return &CityDBRepository{}
}

func (r *CityDBRepository) FindAll(ctx context.Context, params *CityPageParams) (CityPageResponse, error) {
	return sqlDB.NewPageQuery[CityResponse](
		ctx,
		types.NewPageRequest(params.Page, params.Size, []types.Sort{types.NewSort(types.ASC, "name")}),
		findAllquery,
		params.Name,
		params.UF,
	).Execute()
}

func (r *CityDBRepository) FindById(ctx context.Context, id uuid.UUID) (*CityResponse, error) {
	return sqlDB.NewCachedQuery[CityResponse](ctx,
		cacheDB.NewCache[CityResponse](fmt.Sprintf(cityCachePrefix, id), 10*time.Second),
		findByIdQuery,
		id,
	).One()
}

func (r *CityDBRepository) ExistsById(ctx context.Context, id uuid.UUID) (*bool, error) {
	return sqlDB.NewQuery[bool](ctx, existsByIdQuery, id).One()
}

func (r *CityDBRepository) ExistsByUniqueKey(ctx context.Context, name string, uf string) (*bool, error) {
	return sqlDB.NewQuery[bool](ctx, existsByUniqueKeyQuery, name, uf).One()
}

func (r *CityDBRepository) Insert(ctx context.Context, model *CityCreateRequest) (*CityResponse, error) {
	return sqlDB.NewQuery[CityResponse](ctx, insertQuery, model.Name, model.UF).One()
}

func (r *CityDBRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return sqlDB.NewStatement(ctx, deleteQuery, id).Execute()
}
