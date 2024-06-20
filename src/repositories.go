package src

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
}

type CityDBRepository struct{}

func NewCityDBRepository() *CityDBRepository {
	return &CityDBRepository{}
}
