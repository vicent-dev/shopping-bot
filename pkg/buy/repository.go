package buy

var (
	r Repository
)

type Repository interface {
	Get(int64) *Cart
	Store(*Cart)
}

func getRepository() Repository {
	if r != nil {
		return r
	}

	r = NewJsonRepository()

	return r
}
