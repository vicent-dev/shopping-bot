package buy

import (
	"shopping-bot/pkg/storage"
)

type JsonRepository struct {
	carts map[int64]*Cart
}

func (j JsonRepository) Store(cart *Cart) {
	if j.carts == nil {
		j.carts = make(map[int64]*Cart)
	}

	j.carts[cart.Channel] = cart
	storage.Store[map[int64]*Cart](j.carts)
}

func NewJsonRepository() JsonRepository {
	return JsonRepository{}
}

func (j JsonRepository) Get(c int64) *Cart {
	if j.carts == nil {
		j.carts = make(map[int64]*Cart)
	}

	storage.Load[map[int64]*Cart](j.carts)

	if cart, ok := j.carts[c]; ok {
		return cart
	}

	j.carts[c] = &Cart{
		Channel: c,
	}

	return j.carts[c]
}
